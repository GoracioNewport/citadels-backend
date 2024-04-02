package game

type CardDto struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Color       CardColorDto `json:"color"`
	Description []string     `json:"description"`
	Price       int          `json:"price"`
	Image       string       `json:"image"`
	Active      bool         `json:"active"`
}
