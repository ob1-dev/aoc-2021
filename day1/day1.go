package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() {
	f, err := os.Open("input.txt")
	check((err))

	scanner := bufio.NewScanner(f)

	var last_line string
	line_counter := 0
	increase_counter := 0

	for scanner.Scan() {
		// if line_counter == 0 {
		// 	continue
		// }
		next_line := scanner.Text()
		if next_line > last_line {
			increase_counter = increase_counter + 1
		}
		last_line = next_line
		line_counter = line_counter + 1
	}

	fmt.Println(increase_counter)

	fmt.Println("Hello, World!")
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func part2() {
	f, err := os.Open("input2.txt")
	check((err))

	scanner := bufio.NewScanner(f)

	var last_four []int

	increased_count := 0

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		last_four = append(last_four, val)

		if len(last_four) >= 4 {
			if len(last_four) > 4 {
				last_four = RemoveIndex(last_four, 0)
			}
			if sum(last_four[:3]) < sum(last_four[1:]) {
				increased_count += 1
			}
		}
	}
	fmt.Println(increased_count)
}

func main() {

	part2()
}
