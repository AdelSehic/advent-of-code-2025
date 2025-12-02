package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/AdelSehic/advent-of-code-2025/helpers"
)

func part1() int {
	buff := helpers.NewBuffer[int]()

	for i := range 100 {
		buff.Enqueue(&i)
	}
	buff.SetIndex(50)

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		dir := string(str[0])
		move, err := strconv.Atoi(str[1:])
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		if dir == "R" {
			buff.MoveRight(move)
		} else {
			buff.MoveLeft(move)
		}
		if *buff.Get() == 0 {
			counter++
		}
	}
	return counter
}

func part2() int {
	buff := helpers.NewBuffer[int]()

	for i := range 100 {
		buff.Enqueue(&i)
	}
	buff.SetIndex(50)

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		dir := string(str[0])
		move, err := strconv.Atoi(str[1:])
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		moveFunc := buff.Next
		if dir == "L" {
			moveFunc = buff.Previous
		}
		for range move {
			moveFunc()
			if *buff.Get() == 0 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
