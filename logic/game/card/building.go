package card

import (
	"citadels-backend/logic/game/card/template"
	gameDto "citadels-backend/ws/dto/game"
)

type Building struct {
	Name        string
	Color       template.BuildingColor
	Description []string
	Price       int
	Image       string
	Id          int
}

func (b *Building) ToDto() gameDto.CardDto {
	return gameDto.CardDto{
		Id:          b.Id,
		Name:        b.Name,
		Color:       b.Color.ToDto(),
		Description: b.Description,
		Price:       b.Price,
		Image:       b.Image,
	}
}
