package main

import "fmt"

const playerOneSP = 4
const playerTwoSP = 6

var cache = map[[5]int][2]int64{}

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	res := recursePlay([2]int{playerOneSP, playerTwoSP}, [2]int{0, 0}, 0)
	fmt.Printf("Solution 2: Player 1 wins %d universes, Player 2 wins %d\n", res[0], res[1])
}

func recursePlay(pos [2]int, score [2]int, player int) [2]int64 {

	if score[0] >= 21 {
		return [2]int64{1, 0}
	}

	if score[1] >= 21 {
		return [2]int64{0, 1}
	}

	ck := [5]int{pos[0], pos[1], score[0], score[1], player}
	if v, ok := cache[ck]; ok {
		return v
	}

	var universeWins [2]int64

	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				newPos := pos
				newPos[player] = pos[player] + r1 + r2 + r3
				if newPos[player] > 10 {
					newPos[player] = newPos[player] % 10
				}
				newScore := score
				newScore[player] = score[player] + newPos[player]
				nextPlayer := 1
				if player == 1 {
					nextPlayer = 0
				}

				newUniverseWins := recursePlay(newPos, newScore, nextPlayer)
				universeWins[0] += newUniverseWins[0]
				universeWins[1] += newUniverseWins[1]
			}
		}
	}

	cache[ck] = universeWins

	return universeWins
}

func partOne() {

	nbTurns := 0
	nbDiceRolls := 0
	lastDiceRoll := 0
	posPlayerOne := playerOneSP
	scorePlayerOne := 0
	posPlayerTwo := playerTwoSP
	scorePlayerTwo := 0

	for {
		nbTurns++

		//PLAYER ONE
		playerOneRolls := rollThreeTimes(lastDiceRoll)
		nbDiceRolls += 3
		posPlayerOne += sum(playerOneRolls)
		if posPlayerOne > 10 {
			posPlayerOne = posPlayerOne % 10
			if posPlayerOne == 0 {
				posPlayerOne = 10
			}
		}
		scorePlayerOne += posPlayerOne
		if scorePlayerOne >= 1000 {
			break
		}
		//fmt.Printf("Rolls: %+v, pos: %d, score: %d\n", playerOneRolls, posPlayerOne, scorePlayerOne)

		//PLAYER TWO
		playerTwoRolls := rollThreeTimes(playerOneRolls[2])
		nbDiceRolls += 3
		posPlayerTwo += sum(playerTwoRolls)
		if posPlayerTwo > 10 {
			posPlayerTwo = posPlayerTwo % 10
			if posPlayerTwo == 0 {
				posPlayerTwo = 10
			}
		}
		scorePlayerTwo += posPlayerTwo
		if scorePlayerTwo >= 1000 {
			break
		}

		lastDiceRoll = playerTwoRolls[2]
	}

	totalScore := scorePlayerOne * nbDiceRolls
	if scorePlayerOne > scorePlayerTwo {
		totalScore = scorePlayerTwo * nbDiceRolls
	}

	fmt.Printf("Score player 1: %d, Score player 2: %d, Number of Turns: %d, Number of dice rolls : %d, Solution 1: %d\n", scorePlayerOne, scorePlayerTwo, nbTurns, nbDiceRolls, totalScore)
}

func sum(vals []int) int {
	s := 0
	for _, v := range vals {
		s += v
	}

	return s
}

func rollThreeTimes(lastDiceRoll int) []int {
	return []int{lastDiceRoll + 1, lastDiceRoll + 2, lastDiceRoll + 3}
}
