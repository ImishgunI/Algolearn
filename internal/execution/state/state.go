package state

type Step struct {
	Array    []int `json:"array"`
	Active   []int `json:"active"`
	Swapping []int `json:"swapping"`
}
