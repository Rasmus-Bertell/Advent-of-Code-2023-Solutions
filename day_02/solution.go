package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	count int
	color string
}

func main() {
	// input, err := os.ReadFile("input_01.txt.example")
	input, err := os.ReadFile("input_01.txt")

	if err != nil {
		log.Fatal(err)
	}

	problem_01(string(input))

	// input, err = os.ReadFile("input_02.txt.example")
	input, err = os.ReadFile("input_02.txt")

	if err != nil {
		log.Fatal(err)
	}

	problem_02(string(input))
}

func problem_01(input string) {
	log.Println("Problem 01 - Start")

	games := make(map[int][][]Pair)

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		game, moves := parse_line(line)

		games[game] = moves
	}

	valid, _ := find_valid_games(games, []Pair{{12, "red"}, {13, "green"}, {14, "blue"}})

	sum := 0

	for game := range valid {
		sum += game
	}

	log.Println("Valid games sum:", sum)

	log.Println("Problem 01 - End")
}

func problem_02(input string) {
	log.Println("Problem 02 - Start")

	games := make(map[int][][]Pair)

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		game, moves := parse_line(line)

		games[game] = moves
	}

	lowest := find_lowest_cubes(games)

	sum := 0

	for _, value := range lowest {
		power := 1

		for _, count := range value {
			power *= count
		}

		sum += power
	}

	log.Println("Power of games:", sum)

	log.Println("Problem 02 - End")
}

func parse_line(line string) (game int, moves [][]Pair) {
	re, err := regexp.Compile(`^Game (\d+): (.*)$`)

	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindStringSubmatch(line)

	if matches == nil {
		log.Fatal(matches)
	}

	game, err = strconv.Atoi(matches[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, move := range strings.Split(matches[2], ";") {
		var tmp []Pair

		for _, pair := range strings.Split(move, ",") {
			tmp = append(tmp, parse_pair(strings.TrimSpace(pair)))
		}

		moves = append(moves, tmp)
	}

	return
}

func parse_pair(pair string) Pair {
	re, err := regexp.Compile(`^(\d+)\s+(\S+)$`)

	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindStringSubmatch(pair)

	if matches == nil {
		log.Fatal(matches)
	}

	count, err := strconv.Atoi(matches[1])

	if err != nil {
		log.Fatal(err)
	}

	return Pair{count, matches[2]}
}

func find_valid_games(games map[int][][]Pair, constraints []Pair) (valid, invalid map[int][][]Pair) {
	valid = make(map[int][][]Pair)
	invalid = make(map[int][][]Pair)

	for game, moves := range games {
		found := false

		for _, move := range moves {
			for _, pair := range move {
				for _, constraint := range constraints {
					if pair.color == constraint.color && pair.count > constraint.count {
						found = true
					}
				}
			}
		}

		if found {
			invalid[game] = moves
		} else {
			valid[game] = moves
		}
	}

	return
}

func find_lowest_cubes(games map[int][][]Pair) (lowest map[int]map[string]int) {
	lowest = make(map[int]map[string]int)

	for game, moves := range games {
		lowest[game] = make(map[string]int)

		for _, move := range moves {
			for _, pair := range move {
				if pair.count > lowest[game][pair.color] {
					lowest[game][pair.color] = pair.count
				}
			}
		}
	}

	return
}
