package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, NewGame(line))
	}

	sum := 0

	for _, game := range games {
		minSubset := game.MinSubsetToBeValid()
		sum += minSubset.GetPower()
	}

	fmt.Println(sum)
}

func assert(condition bool, msg interface{}) {
	if !condition {
		panic(msg)
	}
}

type Game struct {
	ID      int
	Subsets []Subset
}

type Subset struct {
	Red   int
	Green int
	Blue  int
}

func NewGame(line string) Game {
	game := Game{}
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	assert(scanner.Scan(), "Could not start scanning: "+line)
	assert(scanner.Text() == "Game", "All lines should start with Game: "+line)

	assert(scanner.Scan(), "Could not scan game ID: "+line)
	idString := scanner.Text()
	assert(idString[len(idString)-1] == ':', "Game ID should end with colon: "+line)
	idString = idString[:len(idString)-1]

	id, err := strconv.Atoi(idString)
	assert(err == nil, err)

	game.ID = id

	current := Subset{}

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		assert(err == nil, err)

		assert(scanner.Scan(), "Color not found: "+line)
		color := scanner.Text()
		switch color[0] {
		case 'r':
			current.Red = val
		case 'g':
			current.Green = val
		case 'b':
			current.Blue = val
		default:
			assert(false, "Wrong color: "+color)
		}

		if color[len(color)-1] == ';' {
			game.Subsets = append(game.Subsets, current)
			current = Subset{}
		}
	}

	game.Subsets = append(game.Subsets, current)

	return game
}

func (g *Game) MinSubsetToBeValid() Subset {
	minSubset := Subset{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, s := range g.Subsets {
		minSubset.Red = max(minSubset.Red, s.Red)
		minSubset.Blue = max(minSubset.Blue, s.Blue)
		minSubset.Green = max(minSubset.Green, s.Green)
	}

	return minSubset
}

func (s *Subset) GetPower() int {
	return s.Red * s.Green * s.Blue
}
