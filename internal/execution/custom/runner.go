package custom

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

// Step должен совпадать с тем, что возвращает код пользователя
type Step struct {
	Array    []int `json:"array"`
	Active   []int `json:"active"`
	Swapping []int `json:"swapping"`
}

// RunUserCode принимает код функции Run (на Go) и входной массив, возвращает шаги
func RunUserCode(userCode string, input []int) ([]Step, error) {
	ctx := context.Background()
	cli, err := client.New(client.FromEnv, client.WithHost(os.Getenv("DOCKER_HOST")))
	if err != nil {
		return nil, fmt.Errorf("docker client: %w", err)
	}
	defer cli.Close()

	// Полный код программы с обёрткой
	fullCode := fmt.Sprintf(`
package main

import (
    "encoding/json"
    "os"
    "strconv"
    "strings"
)

type Step struct {
    Array    []int `+"`json:\"array\"`"+`
    Active   []int `+"`json:\"active\"`"+`
    Swapping []int `+"`json:\"swapping\"`"+`
}

%s

func main() {
    inputStr := os.Args[1] // например "5,3,1"
    parts := strings.Split(inputStr, ",")
    input := make([]int, len(parts))
    for i, p := range parts {
        n, err := strconv.Atoi(strings.TrimSpace(p))
        if err != nil {
            panic("invalid input")
        }
        input[i] = n
    }
    steps := Run(input)
    out, err := json.Marshal(steps)
    if err != nil {
        panic("marshal error")
    }
    os.Stdout.Write(out)
}
`, userCode)

	inputStr := intsToString(input)
	config_container := &container.Config{
		Image:      "golang:1.26.0-alpine",
		Cmd:        []string{"sh", "-c", "go run /main.go " + inputStr},
		Tty:        false,
		WorkingDir: "/",
	}
	host_config := &container.HostConfig{
		AutoRemove: false,
		Resources: container.Resources{
			Memory:   8192 * 1024 * 1024, // 8GB
			CPUQuota: 600000,             // 6.0 CPU
		},
		NetworkMode: "none",
	}
	opt := client.ContainerCreateOptions{
		Config:     config_container,
		HostConfig: host_config,
	}
	// Создаём контейнер
	resp, err := cli.ContainerCreate(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("container create: %w", err)
	}

	defer func() {
		// Удаляем контейнер в любом случае
		_, _ = cli.ContainerRemove(ctx, resp.ID, client.ContainerRemoveOptions{Force: true})
	}()

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	hdr := &tar.Header{
		Name: "main.go",
		Mode: 0644,
		Size: int64(len(fullCode)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(fullCode)); err != nil {
		return nil, err
	}
	tw.Close()

	_, err = cli.CopyToContainer(ctx, resp.ID, client.CopyToContainerOptions{
		DestinationPath: "/",
		Content:         &buf,
	})
	if err != nil {
		return nil, fmt.Errorf("copy to container: %w", err)
	}

	start := client.ContainerStartOptions{}
	// Запускаем
	_, err = cli.ContainerStart(ctx, resp.ID, start)
	if err != nil {
		return nil, fmt.Errorf("container start: %w", err)
	}

	timeout := time.After(30 * time.Second)
	for {
		select {
		case <-timeout:
			cli.ContainerStop(ctx, resp.ID, client.ContainerStopOptions{})
			return nil, fmt.Errorf("execution timeout")
		default:
			info, err := cli.ContainerInspect(ctx, resp.ID, client.ContainerInspectOptions{})
			if err != nil {
				return nil, err
			}
			if !info.Container.State.Running {
				goto collectLogs
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
collectLogs:

	// Читаем stdout
	out, err := cli.ContainerLogs(ctx, resp.ID, client.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, fmt.Errorf("logs: %w", err)
	}
	defer out.Close()

	var logBuf bytes.Buffer
	// Docker multiplexes stdout/stderr, нужно демультиплексировать
	_, err = stdcopy.StdCopy(&logBuf, &logBuf, out)
	if err != nil {
		return nil, fmt.Errorf("stdcopy: %w", err)
	}

	var steps []Step
	if err := json.Unmarshal(logBuf.Bytes(), &steps); err != nil {
		return nil, fmt.Errorf("invalid output: %v (output: %s)", err, logBuf.String())
	}
	return steps, nil
}

func intsToString(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	b := strings.Builder{}
	s := fmt.Sprintf("%d", arr[0])
	b.WriteString(s)
	for _, v := range arr[1:] {
		b.WriteByte(',')
		fmt.Fprintf(&b, "%d", v)
	}
	return b.String()
}
