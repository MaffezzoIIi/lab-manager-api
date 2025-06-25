package req

type CreateLabRequest struct {
	Name        string   `json:"name"`
	Local       string   `json:"local"`
	Acessible   bool     `json:"acessible"`
	PcNumbers   int      `json:"pcNumbers"`
	Status      string   `json:"status"`
	Softwares   []string `json:"softwares"`
	Description string   `json:"description,omitempty"`
}
