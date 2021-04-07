package response

type LuckyKResponse struct {
	Date       []string  `json:"date"`
	Total      []float64 `json:"total"`
	Share      []float64 `json:"share"`
	Value      []float64 `json:"value"`
	Percentage []float64 `json:"percentage"`
}
