package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

	valid_numbers := parse_valid_numbers(input)

	sum := 0

	for _, num := range valid_numbers {
		sum += num
	}

	log.Println("Sum of parts:", sum)

	log.Println("Problem 01 - End")
}

func problem_02(input string) {
	log.Println("Problem 02 - Start")

	valid_gears := parse_valid_gears(input)

	sum := 0

	for _, num := range valid_gears {
		sum += (num[0] * num[1])
	}

	log.Println("Gear ratio:", sum)

	log.Println("Problem 02 - End")
}

func parse_valid_numbers(input string) []int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	var prev_line string
	var curr_line string

	var valid_numbers []int

	numbers_re, err := regexp.Compile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	for line_number, line := range lines {
		if line_number == 0 {
			curr_line = line
			continue
		}

		line_numbers := numbers_re.FindAllStringIndex(curr_line, -1)

		for _, number := range line_numbers {
			num, found := extract_valid_number(
				number,
				[]string{prev_line, curr_line, line},
			)

			if found {
				valid_numbers = append(valid_numbers, num)
			}
		}

		prev_line = curr_line
		curr_line = line
	}

	line_numbers := numbers_re.FindAllStringIndex(curr_line, -1)

	for _, number := range line_numbers {
		num, found := extract_valid_number(
			number,
			[]string{prev_line, curr_line},
		)

		if found {
			valid_numbers = append(valid_numbers, num)
		}
	}

	return valid_numbers
}

func parse_valid_gears(input string) [][]int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	var prev_line string
	var curr_line string

	var valid_gears [][]int

	gears_re, err := regexp.Compile(`\*`)

	if err != nil {
		log.Fatal(err)
	}

	for line_number, line := range lines {
		if line_number == 0 {
			curr_line = line
			continue
		}

		possible_gears := gears_re.FindAllStringIndex(curr_line, -1)

		for _, gear := range possible_gears {
			num := extract_valid_gears(gear, []string{prev_line, curr_line, line}, line_number)

			if num != nil {
				valid_gears = append(valid_gears, num)
			}
		}

		prev_line = curr_line
		curr_line = line
	}

	possible_gears := gears_re.FindAllStringIndex(curr_line, -1)

	for _, gear := range possible_gears {
		num := extract_valid_gears(gear, []string{prev_line, curr_line}, 0)

		if num != nil {
			valid_gears = append(valid_gears, num)
		}
	}

	return valid_gears
}

func extract_valid_number(index []int, lines []string) (int, bool) {
	symbols_re, err := regexp.Compile(`[^\d.]+`)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		found := symbols_re.FindStringIndex(line[max(index[0]-1, 0):min(index[1]+1, len(line))])

		if found != nil {
			num, err := strconv.Atoi(lines[1][index[0]:index[1]])

			if err != nil {
				log.Fatal(err)
			}

			return num, true
		}
	}

	return 0, false
}

type number_line struct {
	line   int
	number []int
}

func extract_valid_gears(index []int, lines []string, line_number int) (valid_gears []int) {
	numbers_re, err := regexp.Compile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	found := make([]number_line, 0)

	for line_num, line := range lines {
		if len(line) == 0 {
			continue
		}

		numbers := numbers_re.FindAllStringIndex(line, -1)

		for _, num := range numbers {
			found = append(found, number_line{line_num, num})
		}
	}

	for _, num := range found {
		if (num.number[0] >= index[0] &&
			num.number[0] <= index[1] ||
			num.number[1] >= index[0] &&
				num.number[1] <= index[1]) ||
			(num.number[0] < index[0] &&
				num.number[1] > index[1]) {
			tmp, err := strconv.Atoi(lines[num.line][num.number[0]:num.number[1]])

			if err != nil {
				log.Fatal(err)
			}

			valid_gears = append(valid_gears, tmp)
		}
	}

	if len(valid_gears) != 2 {
		return nil
	}

	return
}
