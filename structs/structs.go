package structs

type PertRequest struct {
	RawData      string `json:"rawData"`
	Perturb      string `json:"perturb"`      // "add_noise", "downsample", "jitter"
	PerturbLevel int    `json:"perturbLevel"` // 1-3
}