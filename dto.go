package blue_green_state_monitor_go

type bgState struct {
	OriginNamespace nsVersion `json:"originNamespace"`
	PeerNamespace   nsVersion `json:"peerNamespace"`
	UpdateTime      string    `json:"updateTime"`
}

type nsVersion struct {
	Name    string  `json:"name"`
	State   string  `json:"state"`
	Version *string `json:"version"`
}

func newVersion(v string) *string {
	if v == "" {
		return nil
	} else {
		return &v
	}
}
