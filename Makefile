build:
	go build -o main cmd/server/main.go

up:
	docker-compose up --build

docker-build:
	docker-compose --build

down:
	docker-compose down -v

clean:
	rm -f main
