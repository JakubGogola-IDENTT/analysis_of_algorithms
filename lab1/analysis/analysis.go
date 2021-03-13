package analysis

import (
	"flag"
	"fmt"
	el "lab1/election"
)

// Analysis struct contains all params required for analysis of algorithm
type Analysis struct {
	lowerLimit int
	upperLimit int
	step       int
	repeats    int
}

// ParseFlags parses input flags with analysis parameters.
func (a *Analysis) ParseFlags() {
	flag.IntVar(&a.lowerLimit, "l", 100, "lower limit")
	flag.IntVar(&a.upperLimit, "u", 10000, "upper limit")
	flag.IntVar(&a.step, "s", 100, "step size")
	flag.IntVar(&a.repeats, "r", 1000, "repeats number")
	flag.Parse()
}

// RunTests runs all tests.
func (a *Analysis) RunTests() {
	a.ParseFlags()
	a.TestScenario2()
	a.TestScenario3()
	a.TestEx4()
}

// TestScenario2 tests second scenario implementation.
func (a *Analysis) TestScenario2() {
	f, w := createFileWithWriter("scenario2_tests.csv")
	defer f.Close()

	_, err := w.WriteString("n,slots\n")
	checkError(err)

	fmt.Println("### Tests for scenario 2 ###")
	for n := a.lowerLimit; n <= a.upperLimit; n += a.step {
		printProgress(n)

		for r := 0; r < a.repeats; r++ {
			slots := el.ElectByScenario2(n)
			_, err := w.WriteString(fmt.Sprintf("%d,%d\n", n, slots))
			checkError(err)
		}

		w.Flush()
	}
	fmt.Println("########################")
}

// TestScenario3 tests third scenario implementation.
func (a *Analysis) TestScenario3() {
	f, w := createFileWithWriter("scenario3_tests.csv")
	defer f.Close()

	_, err := w.WriteString("u,n,slots,rounds\n")
	checkError(err)

	fmt.Println("### Tests for scenario 3 ###")
	for u := a.lowerLimit; u < a.upperLimit/2; u += a.step {
		printProgress(u)

		ns := []int{2, u / 2, u}

		for _, n := range ns {
			for r := 0; r < a.repeats; r++ {
				slots, rounds := el.ElectByScenario3(u, n)
				_, err := w.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", u, n, slots, rounds))
				checkError(err)
			}
		}

		w.Flush()
	}
	fmt.Println("########################")
}

// TestEx4 tests third scenario implementation.
func (a *Analysis) TestEx4() {
	f, w := createFileWithWriter("ex4_tests.csv")
	defer f.Close()

	_, err := w.WriteString("u,n,slots,rounds\n")
	checkError(err)

	fmt.Println("### Tests for exercise 4 ###")
	u := 1000
	printProgress(u)

	for n := 2; n <= u; n++ {
		for r := 0; r < a.repeats; r++ {
			slots, rounds := el.ElectByScenario3(u, n)
			_, err := w.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", u, n, slots, rounds))
			checkError(err)
		}
	}

	w.Flush()

	fmt.Println("########################")
}
