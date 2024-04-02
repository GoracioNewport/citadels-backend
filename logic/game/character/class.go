package character

import "citadels-backend/ws/dto/game"

type Class int

const (
	Architect Class = iota
	Assassin
	Bishop
	King
	Magician
	Merchant
	Thief
	Warlord
)

func (c Class) ToDto() game.CharacterClassDto {
	switch c {
	case Architect:
		return game.Architect
	case Assassin:
		return game.Assassin
	case Bishop:
		return game.Bishop
	case King:
		return game.King
	case Magician:
		return game.Magician
	case Merchant:
		return game.Merchant
	case Thief:
		return game.Thief
	case Warlord:
		return game.Warlord
	}

	return "unknown"
}

func ClassFromString(s string) (class Class, ok bool) {
	switch s {
	case "architect":
		return Architect, true
	case "assassin":
		return Assassin, true
	case "bishop":
		return Bishop, true
	case "king":
		return King, true
	case "magician":
		return Magician, true
	case "merchant":
		return Merchant, true
	case "thief":
		return Thief, true
	case "warlord":
		return Warlord, true
	}

	return 0, false
}
