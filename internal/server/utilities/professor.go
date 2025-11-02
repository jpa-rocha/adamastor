package utilities

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/go-playground/form/v4"
)

const (
	levelOne   = 1
	levelTwo   = 2
	levelThree = 3
	levelFour  = 4
	levelFive  = 5

	numVals = 2
)

type Game struct {
	NumRounds  int
	Level      int
	Win        int
	Loss       int
	Result     int
	Answer     int
	Final      int
	Round      []int
	Op         string
	Operations []string
	Err        error
}

func (g *Game) MakeRound() {
	if g.Op == "R" {
		g.Op = g.Operations[rand.Intn(3)]
	}
	g.numGenerator()
	g.generateResults()
}

func (g *Game) generateResults() {
	switch g.Op {
	case "+":
		g.Result = g.Round[0] + g.Round[1]
	case "-":
		g.Result = g.Round[0] - g.Round[1]
	case "*":
		g.Result = g.Round[0] * g.Round[1]
	case "/":
		g.Result = g.Round[0] / g.Round[1]
	}
}

func (g *Game) numGenerator() {
	var top int
	var bot int
	var topMin int
	var botMin int
	switch g.Level {
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
	g.Round[0] = rand.Intn(top-topMin) + topMin
	g.Round[1] = rand.Intn(bot-botMin) + botMin
	if g.Round[0] <= g.Round[1] {
		g.Round = swap(g.Round)
	}
	if g.Round[1] == 0 {
		g.Round[1]++
	}
}

func (g *Game) PrepareRound(r *http.Request) {
	g.Round = make([]int, numVals)
	err := r.ParseForm()
	if err != nil {
		g.Err = err
		log.Panic(err)
	}
	decoder := form.NewDecoder()

	err = decoder.Decode(&g, r.Form)
	if err != nil {
		g.Err = err
		log.Panic(err)
	}

	log.Println(g)
	g.Operations = []string{"+", "-", "*", "/"}
}

func swap(nums []int) []int {
	nums[0], nums[1] = nums[1], nums[0]
	return nums
}
