package builder

// ManualTrigger this is the POST payload expected
// when a manual trigger happens from the admin panel.
type ManualTrigger struct {
	Who        string `json:"who"`
	ForceBuild bool   `json:"force_build"`
}
