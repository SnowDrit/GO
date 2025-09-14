package game

import "strings"

var player Player
var rooms map[string]Room
var doorLocked bool

func InitGame() {
	rooms = map[string]Room{
		"комната": {
			Description: "ты в своей комнате. можно пройти - коридор",
			Items: []Item{
				{Name: "ключи", Place: "стол"},
				{Name: "конспекты", Place: "стол"},
				{Name: "рюкзак", Place: "стул"},
			},
			Transitions: []string{"коридор"},
		},
		"кухня": {
			Description: "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
			Items: []Item{
				{Name: "чай", Place: "стол"},
			},
			Transitions: []string{"коридор"},
		},
		"коридор": {
			Description: "ничего интересного. можно пройти - кухня, комната, улица",
			Items:       []Item{},
			Transitions: []string{"кухня", "комната", "улица"},
		},
		"улица": {
			Description: "на улице весна. можно пройти - домой",
			Items:       []Item{},
			Transitions: []string{"домой"},
		},
	}
	player = Player{
		CurrentRoom: "кухня",
		Inventory:   []string{},
		HasBackpack: false,
	}
	doorLocked = true
}

func HandleCommand(command string) string {
	parts := strings.Split(command, " ")
	action := parts[0]

	switch action {
	case "осмотреться":
		return handleLook()
	case "идти":
		if len(parts) < 2 {
			return "Некорректная команда"
		}
		return handleGo(parts[1])
	case "взять":
		if len(parts) < 2 {
			return "Некорректная команда"
		}
		return handleTake(parts[1])
	case "надеть":
		if len(parts) < 2 {
			return "Некорректная команда"
		}
		return handleWear(parts[1])
	case "применить":
		if len(parts) < 3 {
			return "Некорректная команда"
		}
		return handleUse(parts[1], parts[2])
	default:
		return "неизвестная команда"
	}
}
