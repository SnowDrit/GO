package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	CurrentRoom string
	Inventory   []string
	HasBackpack bool
}

type Room struct {
	Description string
	Items       []Item
	Transitions []string
}

type Item struct {
	Name  string
	Place string
}

var player Player
var rooms map[string]Room
var doorLocked bool

func initGame() {
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

func handleCommand(command string) string {
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

func handleLook() string {
	currentRoom := rooms[player.CurrentRoom]

	switch player.CurrentRoom {
	case "кухня":
		if player.HasBackpack && len(player.Inventory) > 0 {
			return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		return currentRoom.Description
	case "комната":
		if len(currentRoom.Items) == 0 {
			return "пустая комната. можно пройти - коридор"
		}

		tableItems := []string{}
		chairItems := []string{}

		for _, item := range currentRoom.Items {
			if item.Place == "стол" {
				tableItems = append(tableItems, item.Name)
			} else if item.Place == "стул" {
				chairItems = append(chairItems, item.Name)
			}
		}

		description := ""
		if len(tableItems) > 0 {
			description = "на столе: " + strings.Join(tableItems, ", ")
		}
		if len(chairItems) > 0 {
			if description != "" {
				description += ", "
			}
			description += "на стуле: " + strings.Join(chairItems, ", ")
		}
		description += ". можно пройти - " + strings.Join(currentRoom.Transitions, ", ")
		return description
	default:
		return currentRoom.Description
	}
}

func handleGo(where string) string {
	currentRoom := rooms[player.CurrentRoom]

	for _, transition := range currentRoom.Transitions {
		if transition == where {
			if where == "улица" && doorLocked {
				return "дверь закрыта"
			}

			player.CurrentRoom = where

			if where == "кухня" {
				return "кухня, ничего интересного. можно пройти - коридор"
			}

			return rooms[where].Description
		}
	}

	return "нет пути в " + where
}

func handleTake(what string) string {
	currentRoom := rooms[player.CurrentRoom]

	for i, item := range currentRoom.Items {
		if item.Name == what {
			if what != "рюкзак" && !player.HasBackpack {
				return "некуда класть"
			}

			player.Inventory = append(player.Inventory, what)
			currentRoom.Items = append(currentRoom.Items[:i], currentRoom.Items[i+1:]...)
			rooms[player.CurrentRoom] = currentRoom

			return "предмет добавлен в инвентарь: " + what
		}
	}

	return "нет такого"
}

func handleWear(what string) string {
	if what != "рюкзак" {
		return "неизвестная команда"
	}

	currentRoom := rooms[player.CurrentRoom]

	for i, item := range currentRoom.Items {
		if item.Name == "рюкзак" {
			player.HasBackpack = true
			currentRoom.Items = append(currentRoom.Items[:i], currentRoom.Items[i+1:]...)
			rooms[player.CurrentRoom] = currentRoom
			return "вы надели: рюкзак"
		}
	}

	return "нет такого"
}

func handleUse(what, where string) string {
	hasItem := false
	for _, item := range player.Inventory {
		if item == what {
			hasItem = true
			break
		}
	}

	if !hasItem {
		return "нет предмета в инвентаре - " + what
	}

	if what == "ключи" && where == "дверь" {
		if doorLocked {
			doorLocked = false
			return "дверь открыта"
		}
		return "дверь уже открыта"
	}

	return "не к чему применить"
}

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в текстовую игру!")
	fmt.Println("Введите команду или 'выход' для выхода")

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "выход" {
			break
		}

		response := handleCommand(command)
		fmt.Println(response)
	}
}
