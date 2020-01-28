package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	colors := []string{"♠️", "♥️", "♣️", "♦️"}
	numbers := "A23456789XJQK"
	scanner := bufio.NewScanner(os.Stdin)

	rand.Seed(time.Now().Unix())
	pool := map[int]struct{}{}
	for {
		scanner.Scan()
		scanner.Text()
		for {
			id := rand.Int() % 54

			if _, ok := pool[id]; !ok {
				pool[id] = struct{}{}

				if id == 52 {
					fmt.Print("Small")
				} else if id == 53 {
					fmt.Print("Big")
				} else {
					color := id / 13
					num := id % 13

					fmt.Printf("%s %s", colors[color], string(numbers[num]))
				}
				break
			}
		}
		if len(pool) == 54 {
			fmt.Printf("\nclear")
			pool = map[int]struct{}{}
		}
	}
}
