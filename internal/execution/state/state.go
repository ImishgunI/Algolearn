package state

type Step struct {
	Array     []int `json:"array"`
	Active    []int `json:"active"`
	Swapping  []int `json:"swapping"`
	Lines     []int `json:"lines"`
	Pivot     []int `json:"pivot,omitempty"`
	Partition []int `json:"partition,omitempty"`

	LeftRange   []int `json:"leftRange,omitempty"`
	RightRange  []int `json:"rightRange,omitempty"`
	MergedRange []int `json:"mergedRange,omitempty"`
	Writing     []int `json:"writing,omitempty"`
}
