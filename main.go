package main

import (
	"github.com/starlightromero/monsterslayer/actions"
	"github.com/starlightromero/monsterslayer/interactions"
)

var (
	currentRound int
	rounds       []interactions.Round
)

func main() {
	startGame()

	var winner string

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interactions.PrintGreeting()
}

func executeRound() string {
	var playerAttack int
	var playerHeal int
	var monsterAttack int

	currentRound++
	isSpecialRound := currentRound%3 == 0
	interactions.ShowActions(isSpecialRound)
	playerChoice := interactions.GetPlayerChoice(isSpecialRound)

	switch playerChoice {
	case interactions.Attack:
		playerAttack = actions.AttackMonster(false)
	case interactions.Heal:
		playerHeal = actions.HealPlayer()
	case interactions.SpecialAttack:
		playerAttack = actions.AttackMonster(true)
	}

	_, monsterHealth := actions.GetHealth()

	if monsterHealth <= 0 {
		return "Player"
	}

	monsterAttack = actions.AttackPlayer()
	playerHealth, _ := actions.GetHealth()

	round := interactions.Round{
		Action:        playerChoice,
		PlayerHealth:  playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttack:  playerAttack,
		PlayerHeal:    playerHeal,
		MonsterAttack: monsterAttack,
	}

	round.PrintRoundStats()

	rounds = append(rounds, round)

	if playerHealth <= 0 {
		return "Monster"
	}

	return ""
}

func endGame(winner string) {
	interactions.DeclareWinner(winner)
	interactions.WriteLogFile(&rounds)
}
