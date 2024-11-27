package res

type CreateLabResponse struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Local     string   `json:"local"`
	Acessible bool     `json:"acessible"`
	PcNumbers int      `json:"pcNumbers"`
	Status    string   `json:"status"`
	Softwares []string `json:"softwares"`
}
