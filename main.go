package main

import (
	"battleship/game"
	"fmt"
)

var (
	validUsername = "test"
	validPassword = "12345"
)

func main() {
	// Авторизация
	if !authorize() {
		fmt.Println("Выход из системы. До свидания!")
		return
	}

	// Главное меню
	for {
		fmt.Println("\nДобро пожаловать в 'Морской бой'")
		fmt.Println("1. Начать новую игру")
		fmt.Println("2. Продолжить игру")
		fmt.Println("3. Выход")

		var choice int
		fmt.Print("Введите ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			game.StartNewGame()
		case 2:
			game.ContinueGame()
		case 3:
			fmt.Println("Выход из игры. До свидания!")
			return
		default:
			fmt.Println("Неверный ввод. Попробуйте снова.")
		}
	}
}

// authorize выполняет проверку логина и пароля
func authorize() bool {
	var username, password string
	for attempts := 3; attempts > 0; attempts-- {
		fmt.Println("=== Авторизация ===")
		fmt.Print("Введите логин: ")
		fmt.Scan(&username)
		fmt.Print("Введите пароль: ")
		fmt.Scan(&password)

		if username == validUsername && password == validPassword {
			fmt.Println("Авторизация успешна!")
			return true
		}

		fmt.Printf("Неверный логин или пароль. Попробуйте снова. Осталось попыток: %d\n", attempts-1)
	}

	fmt.Println("Авторизация не удалась. Доступ запрещен.")
	return false
}
