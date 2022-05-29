package bot

type CopypastasStruct struct {
	Copypastas []Copypasta `json:"copypastas"`
}

type Copypasta struct {
	name string `json:"name"`
	data string `json:"data"`
}
