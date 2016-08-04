package randomOrgUtil

type GenerateIntegersParams struct {
  ApiKey      string  `json:"apiKey"`
  N           float64 `json:"n"`
  Min         float64 `json:"min"`
  Max         float64 `json:"max"`
  Replacement bool    `json:"replacement"`
  Base        float64 `json:"base"`
}

func GenerateIntegers(params GenerateIntegersParams) ([]int64, error) {
  return []int64{1}, nil
}
