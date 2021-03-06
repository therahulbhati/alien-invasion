package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/therahulbhati/alien-invasion/pkg/invasion"
)

var numberOfAliens int
var maxIterations int
var inputFile string
var seedValue int64

func init() {
	flag.IntVar(&numberOfAliens, "numberOfAliens", 5, "Number of aliens invading")
	flag.IntVar(&maxIterations, "maxIterations", 10000, "Maximum number of iterations")
	flag.StringVar(&inputFile, "inputFile", "", "World map input file")
	flag.Int64Var(&seedValue, "seedValue", 42, "Seed value for deterministic state")

}

func Start() {
	flag.Parse()
	Execute(inputFile, seedValue, numberOfAliens, maxIterations)
}

func Execute(inputFile string, seedValue int64, numberOfAliens, maxIterations int) {
	if numberOfAliens <= 0 {
		log.Fatal("Number of aliens should be positive integer")
	}

	if maxIterations < 0 {
		log.Fatal("Maximum iteration should >= 0")
	}
	if len(inputFile) == 0 {
		log.Fatal("Please enter input file path")
	}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error while opening file: %s, Error: %v", inputFile, err)
	}
	defer file.Close()

	simulator, err := invasion.NewSimulator(file, numberOfAliens, maxIterations, seedValue)
	if err != nil {
		log.Fatal("Error while initiating new Simulator", err)
	}
	simulator.Simulate()
}
