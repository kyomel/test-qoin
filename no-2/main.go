package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id     int
	dice   []int
	points int
}

func rollDice(numDice int) []int {
	dice := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}
	return dice
}

func evaluateDice(players []Player) {
	for i := 0; i < len(players); i++ {
		newDice := []int{}
		for _, die := range players[i].dice {
			switch die {
			case 6:
				players[i].points++
			case 1:
				nextPlayer := (i + 1) % len(players)
				players[nextPlayer].dice = append(players[nextPlayer].dice, 1)
			default:
				newDice = append(newDice, die)
			}
		}
		players[i].dice = newDice
	}
}

func printStatus(players []Player) {
	fmt.Println("==================")
	for _, player := range players {
		fmt.Printf("Pemain #%d (%d): %v\n", player.id, player.points, player.dice)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu per pemain: ")
	fmt.Scan(&numDice)

	players := make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = Player{id: i + 1, dice: rollDice(numDice)}
	}

	round := 1
	for {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)
		for i := 0; i < len(players); i++ {
			players[i].dice = rollDice(len(players[i].dice))
		}
		printStatus(players)

		evaluateDice(players)
		fmt.Println("Setelah evaluasi:")
		printStatus(players)

		activePlayers := 0
		for _, player := range players {
			if len(player.dice) > 0 {
				activePlayers++
			}
		}
		if activePlayers <= 1 {
			break
		}

		round++
	}

	maxPoints := 0
	winner := -1
	for _, player := range players {
		if player.points > maxPoints {
			maxPoints = player.points
			winner = player.id
		}
	}

	fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", winner)
	fmt.Printf("Game dimenangkan oleh pemain #%d dengan %d poin.\n", winner, maxPoints)
}
