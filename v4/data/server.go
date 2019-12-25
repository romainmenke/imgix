package data

type Server struct {
	Commit string `json:"commit"`
	Status struct {
		Healthy   bool `json:"healthy"`
		ReadOnly  bool `json:"read_only"`
		Tombstone bool `json:"tombstone"`
	} `json:"status"`
	Version string `json:"version"`
}
