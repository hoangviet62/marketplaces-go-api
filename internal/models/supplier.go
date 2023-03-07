package internal

type Supplier struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Mobile      string `json:"mobile"`
}
