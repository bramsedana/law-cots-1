package entity

// Equation represents the formula
type Equation struct {
	First     float64 `json:"first"`
	Second    float64 `json:"second"`
	Operation string  `json:"operation"`
}
