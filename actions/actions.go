package actions

import (
	"math/rand"
	"time"
)

var (
	randSouce            = rand.NewSource(time.Now().UnixNano())
	randGenerator        = rand.New(randSouce)
	currentMonsterHealth = monsterHealth
	currentPlayerHealth  = playerHealth
)

func AttackMonster(isSpecialAttack bool) int {
	minAttack := playerMinAttack
	maxAttak := playerMaxAttack

	if isSpecialAttack {
		minAttack = playserMinSpecialAttack
		maxAttak = playerMaxSpecialAttack
	}

	damage := generateRandBetween(minAttack, maxAttak)
	currentMonsterHealth -= damage

	return damage
}

func HealPlayer() int {
	heal := generateRandBetween(playerMinHeal, playerMaxHeal)

	healthDiff := playerHealth - currentPlayerHealth
	if healthDiff >= currentPlayerHealth {
		currentPlayerHealth += heal
		return heal
	} else {
		currentPlayerHealth = playerHealth
		return healthDiff
	}
}

func AttackPlayer() int {
	damage := generateRandBetween(monsterMinAttack, monsterMaxAttack)
	currentPlayerHealth -= damage
	return damage
}

func GetHealth() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
