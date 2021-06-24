package interactions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Attack        = "ATTACK"
	Heal          = "HEAL"
	SpecialAttack = "SPECIAL_ATTACK"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func GetPlayerChoice(hasSpecialAttack bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		switch playerChoice {
		case "1":
			return Attack
		case "2":
			return Heal
		case "3":
			return SpecialAttack
		}
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	playerInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	playerInput = strings.ReplaceAll(playerInput, "\n", "")

	return playerInput, nil
}
