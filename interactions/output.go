package interactions

import (
	"fmt"
	"log"
	"os"
)

type Round struct {
	Action        string
	PlayerAttack  int
	PlayerHeal    int
	MonsterAttack int
	PlayerHealth  int
	MonsterHealth int
}

func (r *Round) PrintRoundStats() {
	switch r.Action {
	case Attack:
		fmt.Printf("Player attacked monster for %d damage.\n", r.PlayerAttack)
	case SpecialAttack:
		fmt.Printf("Player performed a special attack. Monster took %d damage.\n", r.PlayerAttack)
	case Heal:
		fmt.Printf("Player healed. Gained %d points.\n", r.PlayerHeal)
	}

	fmt.Printf("Monster attacked player for %d damage.\n", r.MonsterAttack)
	fmt.Printf("Player health: %d\n", r.PlayerHealth)
	fmt.Printf("Monster health: %d\n", r.MonsterHealth)
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("May the spirit of the ancestors guide your journey")
}

func ShowActions(hasSpecialAttack bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-----------------------------")
	fmt.Println("1) Attack Monster")
	fmt.Println("2) Heal")

	if hasSpecialAttack {
		fmt.Println("3) Special Attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("-----------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("-----------------------------")
	fmt.Printf("%s won!\n", winner)
}

func WriteLogFile(rounds *[]Round) {
	file, err := os.Create("gamelog.txt")
	if err != nil {
		log.Fatal("Saving log file failed. Exiting...")
		return
	}

	for i, v := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(i + 1),
			"Action":                v.Action,
			"Player Attack Damage":  fmt.Sprint(v.PlayerAttack),
			"Player Heal Value":     fmt.Sprint(v.PlayerHeal),
			"Monster Attack Damage": fmt.Sprint(v.MonsterAttack),
			"Player Health":         fmt.Sprint(v.PlayerHealth),
			"Monster Health":        fmt.Sprint(v.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)
		if err != nil {
			log.Fatal("Writing to log file failed. Exiting...")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log.")
}
