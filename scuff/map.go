package scuff

type scuff struct {
	Config scuffConfig `json:"Of"`
	AsMap  map[string]interface{}
}
type scuffConfig struct {
	Location string   `json:"location"`
	Delim    delim    `json:"delim"`
	In       string   `json:"in"`
	Out      string   `json:"out"`
	Rewrite  []string `json:"rewrite"`
}
type delim struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}
