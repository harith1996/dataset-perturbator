package structs

type PertRequest struct {
	ID           string `json:"id"` // UUID
	RawData      string `json:"rawData"`
	Perturb      string `json:"perturb"`      // "addNoise", "downsample", "jitter"
	PerturbLevel int    `json:"perturbLevel"` // 1-3
}
