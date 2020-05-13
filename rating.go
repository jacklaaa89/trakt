package trakt

type Rating struct {
	Score        float64          `json:"rating"`
	Votes        int64            `json:"votes"`
	Distribution map[string]int64 `json:"distribution"`
}
