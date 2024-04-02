package card

import (
	"citadels-backend/logic/game/card/template"
	gameDto "citadels-backend/ws/dto/game"
	"strconv"
)

type Building struct {
	Name        string
	Color       template.BuildingColor
	Description []string
	Price       int
	Image       string
	Id          int
}

var id = 0

func BuildingFromTemplate(template template.BuildingTemplate) Building {
	id++

	var image string
	if template.Image == "" {
		image = string(template.Color.ToDto()) + "_" + strconv.Itoa(template.Price) + ".png"
	} else {
		image = template.Image
	}

	return Building{
		Name:        template.Name,
		Color:       template.Color,
		Description: template.Description,
		Price:       template.Price,
		Image:       image,
		Id:          id - 1,
	}
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
