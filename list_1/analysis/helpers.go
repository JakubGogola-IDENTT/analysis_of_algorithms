package analysis

import (
	"fmt"
	"log"
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
	variance := float64(getTotal(data)) / float64(iterationsCount)

	fmt.Println(variance)
}
