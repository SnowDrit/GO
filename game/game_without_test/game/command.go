package game

import "strings"

func handleLook() string {
	currentRoom := rooms[player.CurrentRoom]

	switch player.CurrentRoom {
	case "кухня":
		if player.HasBackpack && len(player.Inventory) > 0 {
			return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
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
			if what == "рюкзак" {
				return "рюкзак нужно надеть, а не взять. Используйте команду 'надеть рюкзак'"
			}
			if !player.HasBackpack {
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
