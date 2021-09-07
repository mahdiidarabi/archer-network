package main

import (
	"context"
	"fmt"
	"github.com/mahdiidarabi/archer-network/domain"
	"os"
	"strconv"
	"time"
)

func main()  {
	args := os.Args

	totalNumberOfArchersStr := args[1]
	totalNumberOfArchers, err := strconv.Atoi(totalNumberOfArchersStr)
	if err != nil {
		fmt.Println(err)
	}

	archers := make([]domain.Archer, 0, totalNumberOfArchers)

	// create 10 archers
	for i := 0; i < totalNumberOfArchers; i++ {
		newArcher := domain.NewArcher(uint64(totalNumberOfArchers), uint64(i))

		archers = append(archers, newArcher)
	}

	// add the left neighbors
	for i := 1; i < totalNumberOfArchers; i++ {
		archers[i].AddLeftNeighbor(archers[i - 1])
	}

	// add the right neighbors
	for i := 0; i < totalNumberOfArchers - 1; i++ {
		archers[i].AddRightNeighbor(archers[i + 1])
	}

	msgStr := "fire-1"

	err = archers[0].HearFromNeighborsForMessage(context.Background(), msgStr)
	if err != nil {
		fmt.Println(err)
	}

	sleepSeconds := time.Duration(2)*time.Second
	time.Sleep(sleepSeconds)
}