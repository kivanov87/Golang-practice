package main

import (
	"fmt"
	"math"
	"strings"
)

type Evaluator struct {
	Name  string
	Score float64
}

func main() {
	var actorName string
	var academyPoints, evaluatorScore float64
	var evaluatorCount int
	var evaluatorName string

	fmt.Scan(&actorName)
	fmt.Scan(&academyPoints)
	fmt.Scan(&evaluatorCount)

	evaluators := make([]Evaluator, evaluatorCount)

	for i := 0; i < evaluatorCount; i++ {
		fmt.Scan(&evaluatorName)
		fmt.Scan(&evaluatorScore)
		evaluators[i] = Evaluator{Name: evaluatorName, Score: evaluatorScore}
	}

	totalPoints := academyPoints

	for _, evaluator := range evaluators {
		totalPoints += float64(len(strings.ReplaceAll(evaluator.Name, " ", ""))) * evaluator.Score / 2
		if totalPoints > 1250.5 {
			fmt.Printf("Congratulations, %s got a nominee for leading role with %.1f!\n", actorName, totalPoints)
			return
		}
	}

	fmt.Printf("Sorry, %s you need %.1f more!\n", actorName, math.Ceil(1250.5-totalPoints))
}
