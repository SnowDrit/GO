package main

import (
	"bufio"
	"fmt"
	"game"
	"os"
	"strings"
)

func main() {
	game.InitGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в текстовую игру!")
	fmt.Println("")
	fmt.Println("Доступные команды:")
	fmt.Println("  осмотреться - осмотреть текущую комнату")
	fmt.Println("  идти [комната] - переместиться в другую комнату")
	fmt.Println("  взять [предмет] - взять предмет в инвентарь")
	fmt.Println("  надеть [предмет] - надеть предмет (рюкзак)")
	fmt.Println("  применить [предмет] [объект] - использовать предмет на объекте")
	fmt.Println("  выход - выйти из игры")
	fmt.Println("")
	fmt.Println("Доступные комнаты: кухня, комната, коридор, улица")
	fmt.Println("Доступные предметы: чай, ключи, конспекты, рюкзак")
	fmt.Println("")
	fmt.Println("Для начала введите 'осмотреться' чтобы осмотреть текущую комнату")
	fmt.Println("")

	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "выход" {
			break
		}

		response := game.HandleCommand(command)
		fmt.Println(response)
	}
}
