package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	points          int
	amount          int
	winning_numbers []int
	numbers         []int
}

func main() {
	// input, err := os.ReadFile("input_01.txt.example")
	input, err := os.ReadFile("input_01.txt")

	if err != nil {
		log.Fatal(err)
	}

	problem_01(string(input))

	log.Println("")

	// input, err = os.ReadFile("input_02.txt.example")
	input, err = os.ReadFile("input_02.txt")

	if err != nil {
		log.Fatal(err)
	}

	problem_02(string(input))
}

func problem_01(input string) {
	log.Println("Problem 01 - Start")

	cards := parse_cards(input)
	points := calculate_points(cards)

	log.Println("Scratchcard points:", points)

	log.Println("Problem 01 - End")
}

func problem_02(input string) {
	log.Println("Problem 02 - Start")

	cards := parse_cards(input)
	cards = calculate_card_points(cards)
	amount := duplicate_cards(cards)

	log.Println("Scratchcards:", amount)

	log.Println("Problem 02 - End")
}

func parse_cards(input string) (cards []Card) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	for _, line := range lines {
		cards = append(cards, parse_card(line))
	}
	return
}

func parse_card(input string) (card Card) {
	re_number_groups, err := regexp.Compile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	card.amount = 1

	for i, num := range re_number_groups.FindAllString(input, -1) {
		if i == 0 {
			continue
		}

		number, err := strconv.Atoi(num)

		if err != nil {
			log.Fatal(err)
		}

		if i < 11 {
			card.winning_numbers = append(card.winning_numbers, number)
		} else {
			card.numbers = append(card.numbers, number)
		}
	}

	return
}

func calculate_points(cards []Card) (total int) {
	for _, card := range cards {
		for _, num := range card.numbers {
			for _, win := range card.winning_numbers {
				if num == win {
					if card.points == 0 {
						card.points = 1

						continue
					}

					card.points *= 2
				}
			}
		}

		total += card.points
	}

	return
}

func calculate_card_points(cards []Card) (new_cards []Card) {
	new_cards = cards

	for i, card := range cards {
		for _, num := range card.numbers {
			for _, win := range card.winning_numbers {
				if num == win {
					new_cards[i].points++
				}
			}
		}
	}

	return
}
func duplicate_cards(cards []Card) (total int) {
	for i, card := range cards {
		for amount := 0; amount < card.amount; amount++ {
			for duplicate := 1; duplicate < card.points+1; duplicate++ {
				cards[i+duplicate].amount++
			}
		}

		total += card.amount
	}

	return
}
