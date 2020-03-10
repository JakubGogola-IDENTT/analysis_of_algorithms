package analysis

import (
	"fmt"
	"list_1/election"
	"log"
	"math"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func createHistogram(data map[int]int, title string) {
	p, err := plot.New()

	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = title

	values := make(plotter.XYs, len(data))

	i := 0
	for k, v := range data {
		values[i] = plotter.XY{X: float64(k), Y: float64(v)}
		i++
	}

	h, err := plotter.NewHistogram(values, len(data))

	if err != nil {
		log.Fatal(err)
	}

	p.Add(h)
	h.DataRange()

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("hist_%s.png", strings.ReplaceAll(title, " ", "_"))); err != nil {
		log.Fatal(err)
	}

}

func extractValues(data map[int]int) (values []int) {
	for _, v := range data {
		values = append(values, v)
	}

	return values
}

func getTotal(data map[int]int) (total int) {
	for k, v := range data {
		total += k * v
	}

	return total
}

func computeStats(data map[int]int, iterationsCount int) {
	expectedValue := float64(getTotal(data)) / float64(iterationsCount)

	var variance float64

	for k, v := range data {
		variance += float64(v) * math.Pow(float64(k)-expectedValue, 2.0)
	}

	variance = variance / float64(iterationsCount)

	fmt.Printf("Expected value: %f\n", expectedValue)
	fmt.Printf("Variance: %f\n", variance)
}

func computerFirstRoundSuccessProbability(upperLimit, nodesCount, iterationsCount int) {
	firstRoundSuccessCount := 0

	for i := 0; i < iterationsCount; i++ {
		_, roundsCount := election.WithUpperLimit(upperLimit, nodesCount)

		if roundsCount == 1 {
			firstRoundSuccessCount++
		}
	}

	firstRoundSuccessProb := float64(firstRoundSuccessCount) / float64(iterationsCount)

	fmt.Printf("Probability of success in first round: %f\n\n", firstRoundSuccessProb)
}
