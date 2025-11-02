// nolint: forbidigo, gosec, nolintlint
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"time"
	"unicode"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"

	levelOne   = 1
	levelTwo   = 2
	levelThree = 3
	levelFour  = 4
	levelFive  = 5

	numVals   = 2
	numRounds = 5
)

type Game struct {
	numRounds  int
	level      int
	win        int
	loss       int
	result     int
	answer     int
	round      []int
	op         string
	operations []string
	roundStart time.Time
	roundTime  time.Duration
	totalTime  time.Duration
	err        error
}

var ErrInvalidChar = errors.New("error: invalid char detected")

func main() {
	for {
		game := setup(numRounds)
		if game.err != nil {
			return
		}
		RunGame(&game)
		fmt.Println("-------------------------------------------")
		fmt.Println(
			"Average round duration (in seconds):",
			(int(game.totalTime.Seconds()) / game.numRounds),
		)
		fmt.Println("Total game time:", game.totalTime)
		fmt.Println("-------------------------------------------")
		fmt.Printf("Would you like to play again <y> or <n>: ")
		newGame, err := restartGame()
		if err != nil {
			return
		}
		if !newGame {
			fmt.Println("Goodbye!")
			return
		}
	}
}

func setup(numRounds int) Game {
	var game Game
	game.numRounds = numRounds
	game.operations = []string{"+", "-", "*", "/"}
	fmt.Printf("Enter a number between 1 and 5: ")
	game.level, game.err = getNum()
	if game.err != nil {
		fmt.Println(game.err)
		return game
	}
	if game.level > 5 || game.level < 1 {
		fmt.Println("The level should be between 1 and 5.")
		return game
	}
	fmt.Printf("Enter an operation (+, -, *, /): ")

	game.op, game.err = getString()
	if game.err != nil {
		fmt.Println(game.err)
		return game
	}
	if len(game.op) > 1 {
		fmt.Println("You should only enter one char.")
		return game
	}

	if !slices.Contains(game.operations, game.op) {
		fmt.Println("You entered an invalid operation")
		return game
	}
	return game
}

func RunGame(game *Game) {
	change := true
	for {
		if change {
			game.roundStart = time.Now()
			game.round = numGenerator(game.level)
		}
		generateResults(game)
		getAnswer(game)
		if game.answer == game.result {
			game.win++
			change = true
		} else {
			game.loss++
			change = false
		}
		game.roundTime = time.Since(game.roundStart)
		game.totalTime += game.roundTime
		printResults(*game)
		if game.win >= game.numRounds || game.loss >= game.numRounds {
			break
		}
	}
}

func restartGame() (bool, error) {
	for {
		newGame, err := getString()
		if err != nil {
			return false, err
		}
		switch newGame {
		case "y":
			return true, nil
		case "n":
			return false, nil
		default:
			fmt.Println("Invalid char, please enter <y> or <n>.")
		}
	}
}

func getAnswer(game *Game) {
	for {
		fmt.Printf("%d %s %d = ", game.round[0], game.op, game.round[1])
		game.answer, game.err = getNum()
		if game.err == nil {
			break
		} else if errors.Is(game.err, ErrInvalidChar) {
			fmt.Println("invalid char, please enter a number")
		}
	}
}

func generateResults(game *Game) {
	switch game.op {
	case "+":
		game.result = game.round[0] + game.round[1]
	case "-":
		game.result = game.round[0] - game.round[1]
	case "*":
		game.result = game.round[0] * game.round[1]
	case "/":
		game.result = game.round[0] / game.round[1]
	}
}

func printResults(game Game) {
	fmt.Println("SCORE:")
	fmt.Printf("%s", Green)
	for i := 0; i < game.win; i++ {
		fmt.Printf("O")
	}
	fmt.Printf("%s", Reset)
	if game.win < game.numRounds {
		for i := (game.numRounds - game.win); i > 0; i-- {
			fmt.Printf(" ")
		}
	}
	fmt.Printf(" | ")

	fmt.Printf("%s", Red)
	for i := 0; i < game.loss; i++ {
		fmt.Printf("X")
	}
	fmt.Printf("%s", Reset)
	fmt.Printf("\n")
	fmt.Printf("Round lasted: %02s\n", game.roundTime)
}

func getString() (string, error) {
	var level string
	_, err := fmt.Scan(&level)
	if err != nil {
		return "", err
	}
	return level, nil
}

func getNum() (int, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(input); i++ {
		if !unicode.IsDigit(rune(input[i])) {
			return 0, ErrInvalidChar
		}
	}
	numLevel, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return numLevel, nil
}

func numGenerator(level int) []int {
	nums := make([]int, numVals)
	var top int
	var bot int
	var topMin int
	var botMin int
	switch level {
	case levelOne:
		top = 10
		topMin = 1
		bot = top
		botMin = topMin
	case levelTwo:
		top = 100
		topMin = 10
		bot = 10
		botMin = 1
	case levelThree:
		top = 100
		topMin = 10
		bot = top
		botMin = topMin
	case levelFour:
		top = 1000
		topMin = 100
		bot = 100
		botMin = 10
	case levelFive:
		top = 1000
		topMin = 100
		bot = 1000
		botMin = topMin
	}
	nums[0] = rand.Intn(top-topMin) + topMin
	nums[1] = rand.Intn(bot-botMin) + botMin
	if nums[0] < nums[1] {
		nums = swap(nums)
	}
	if nums[1] == 0 {
		nums[1]++
	}
	return nums
}

func swap(nums []int) []int {
	nums[0], nums[1] = nums[1], nums[0]
	return nums
}
