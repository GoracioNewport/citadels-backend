package template

var Factory = BuildingTemplate{
	Name:  "Фабрика",
	Color: Purple,
	Description: []string{
		"Для вас цена постройки других бонусных кварталов уменьшается на 1 золотой.",
	},
	Price: 5,
	Image: "",
}

var Necropolis = BuildingTemplate{
	Name:  "Некрополь",
	Color: Purple,
	Description: []string{
		"Вместо того чтобы оплачивать стоимость некрополя, вы можете построить его, разрушив 1 из кварталов в вашем городе.",
	},
	Price: 5,
	Image: "",
}

var CardCollection = BuildingTemplate{
	Name:  "Коллекция карт",
	Color: Purple,
	Description: []string{
		"В конце игры вы получаете 1 дополнительное очко за каждую карту у вас на руке",
	},
	Price: 5,
	Image: "",
}

var GoldMine = BuildingTemplate{
	Name:  "Золотой рудник",
	Color: Purple,
	Description: []string{
		"Если вы решаете при сборе ресурсов получать золотые из банка, получите 1 дополнительный золотой.",
	},
	Price: 6,
	Image: "",
}

var DragonGate = BuildingTemplate{
	Name:  "Врата дракона",
	Color: Purple,
	Description: []string{
		"В конце игры получите 2 дополнительных очка",
	},
	Price: 6,
	Image: "",
}

var Observatory = BuildingTemplate{
	Name:  "Обсерватория",
	Color: Purple,
	Description: []string{
		"Если вы решаете при сборе ресурсов брать карты, берите 3 карты вместо 2",
	},
	Price: 4,
	Image: "",
}

var Almshouse = BuildingTemplate{
	Name:  "Богадельня",
	Color: Purple,
	Description: []string{
		"Если в конце вашего хода у вас в казне нет золота, получите 1 золотой из банка.",
	},
	Price: 4,
	Image: "",
}

var Basilica = BuildingTemplate{
	Name:  "Базилика",
	Color: Purple,
	Description: []string{
		"В конце игры получите 1 дополнительное очко за каждый квартал в вашем городе, который стоит нечётное число золотых.",
	},
	Price: 4,
	Image: "",
}

var Forge = BuildingTemplate{
	Name:  "Кузница",
	Color: Purple,
	Description: []string{
		"Раз в свой ход вы можете заплатить 2 золотых в банк, чтобы взять на руку 3 карты из колоды кварталов.",
	},
	Price: 5,
	Image: "",
}

var GhostTown = BuildingTemplate{
	Name:  "Квартал призраков",
	Color: Purple,
	Description: []string{
		"В конце игры квартал призраков считается кварталом любого вида по вашему выбору.",
	},
	Price: 2,
	Image: "",
}

var TowerOfCalm = BuildingTemplate{
	Name:  "Башня спокойствия",
	Color: Purple,
	Description: []string{
		"Если в конце игры Башня спокойствия - единственный бонусный квартал в вашем городе, получите 5 дополнительных очков.",
	},
	Price: 5,
	Image: "",
}

var Capitol = BuildingTemplate{
	Name:  "Капитолий",
	Color: Purple,
	Description: []string{
		"Если в конце игры у вас есть хотя бы 3 квартала одного вида, вы получаете 3 дополнительных очка",
	},
	Price: 5,
	Image: "",
}

var Stable = BuildingTemplate{
	Name:  "Конюшня",
	Color: Purple,
	Description: []string{
		"На конюшню не влияет ограничение на количество кварталов, которые мы можете построить за ход.",
	},
	Price: 2,
	Image: "",
}

var SecretVault = BuildingTemplate{
	Name:  "Тайное хранилище",
	Color: Purple,
	Description: []string{
		"Тайное хранилище нельзя построить. В конце игры можете раскрыть карту 'Тайное хранилище' с руки, чтобы получить 3 дополнительных очка.",
	},
	Price: 0,
	Image: "",
}

var Monument = BuildingTemplate{
	Name:  "Монумент",
	Color: Purple,
	Description: []string{
		"Нельзя построить монумент, если в вашем городе 5 или больше кварталов. В отношении того, когда считать город достроенным, монумент считается за 2 квартала.",
	},
	Price: 4,
	Image: "",
}

var Laboratory = BuildingTemplate{
	Name:  "Лаборатория",
	Color: Purple,
	Description: []string{
		"Раз в свой ход вы можете сбросить 1 карту квартала с руки, чтобы взять 2 золотых из банка.",
	},
	Price: 5,
	Image: "",
}

var Museum = BuildingTemplate{
	Name:  "Музей",
	Color: Purple,
	Description: []string{
		"Раз в свой ход вы можете положить под музей 1 карту с руки рубашкой вверх. В конце игры вы получаете 1 дополнительное очко за каждую карту под музеем.",
	},
	Price: 4,
	Image: "",
}

var Fort = BuildingTemplate{
	Name:  "Форт",
	Color: Purple,
	Description: []string{
		"Персонажи 8-го ранга не могут применять своё особое свойство против форта.",
	},
	Price: 3,
	Image: "",
}

var ImperialTreasury = BuildingTemplate{
	Name:  "Имперская казна",
	Color: Purple,
	Description: []string{
		"В конце игры вы получаете 1 дополнительное очко за каждый золотой в вашей казне.",
	},
	Price: 5,
	Image: "",
}

var WellOfWishes = BuildingTemplate{
	Name:  "Колодец желаний",
	Color: Purple,
	Description: []string{
		"В конце игры вы получаете 1 дополнительное очко за каждый бонусный квартал в вашем городе (включая 'Колодец желаний').",
	},
	Price: 5,
	Image: "",
}

var DenOfThieves = BuildingTemplate{
	Name:  "Логово воров",
	Color: Purple,
	Description: []string{
		"Вы можете частично или полностью оплатить стоимость логова воров картами с руки вместо золота (по курсу '1 карта = 1 золотой').",
	},
	Price: 6,
	Image: "",
}

var Memorial = BuildingTemplate{
	Name:  "Памятник",
	Color: Purple,
	Description: []string{
		"Если в конце игры у вас есть корона, получите 5 дополнительных очков.",
	},
	Price: 3,
	Image: "",
}

var GreatWall = BuildingTemplate{
	Name:  "Великая стена",
	Color: Purple,
	Description: []string{
		"Чтобы применить своё особое свойство против любого другого вашего квартала, персонаж 8-го ранга должен заплатить на 1 золотой больше.",
	},
	Price: 6,
	Image: "",
}

var Park = BuildingTemplate{
	Name:  "Парк",
	Color: Purple,
	Description: []string{
		"Если в конце вашего хода у вас на руке нет карт, возьмите 2 карты из колоды кварталов",
	},
	Price: 6,
	Image: "",
}

var Arsenal = BuildingTemplate{
	Name:  "Арсенал",
	Color: Purple,
	Description: []string{
		"В свой ход вы можете разрушить арсенал, чтобы разрушить любой другой квартал по вашему выбору. Нельзя разрушать кварталы в достроенном городе.",
	},
	Price: 3,
	Image: "",
}

var Frame = BuildingTemplate{
	Name:  "Каркас",
	Color: Purple,
	Description: []string{
		"Вместо того чтобы оплачивать стоимость квартала, вы можете построить его, разрушив каркас.",
	},
	Price: 3,
	Image: "",
}

var Theater = BuildingTemplate{
	Name:  "Театр",
	Color: Purple,
	Description: []string{
		"В конце каждого шана выбора вы можете обменять вашу выбранную карту персонажа на карту персонажа другого игрока.",
	},
	Price: 6,
	Image: "",
}

var MagicSchool = BuildingTemplate{
	Name:  "Школа магии",
	Color: Purple,
	Description: []string{
		"Когда вы применяете свойства, позволяющие получить ресурсы за квараталы, школа магии считается кварталом любого вида по вашему выбору",
	},
	Price: 6,
	Image: "",
}

var Quarry = BuildingTemplate{
	Name:  "Каменоломня",
	Color: Purple,
	Description: []string{
		"Вы можете строить кварталы, которые уже есть в вашем городе",
	},
	Price: 5,
	Image: "",
}

var Library = BuildingTemplate{
	Name:  "Библиотека",
	Color: Purple,
	Description: []string{
		"Если вы решаете при сборе ресурсов брать карты, оставьте на руке все взятые карты",
	},
	Price: 6,
	Image: "",
}
