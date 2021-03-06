package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/minskylab/calab"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/spaces/board"
	"github.com/minskylab/calab/systems/cyclic"
)

func main() {

	var c0, _ = colorful.Hex("#1e2031")
	var c1, _ = colorful.Hex("#fbe3a1")

	width, height := 256, 256

	palette := calab.Palette{0: c1, 1: c0, 2: c1, 3: c0}

	// creating the space.
	nh := board.MooreNeighborhood(1, false)
	bound := board.ToroidBounded()
	space := board.MustNew(width, height, nh, bound, board.RandomInit, board.UniformNoise(len(palette)))

	// creating the rule.
	// rule := lifelike.MustNew(lifelike.DayAndNight)
	rule := cyclic.MustNewRockPaperScissor(len(palette), 1, 4)

	// bulk into dynamical system.
	system := calab.BulkDynamicalSystem(space, rule)

	// srv := remote.NewBinaryRemote(3000, "/", pd.binaryChannel)

	vm := calab.NewVM(system)

	pd := experiments.NewPetriDish(vm, palette, 10000)

	startTime := time.Now()

	pd.Run(10 * time.Second)

	fmt.Printf("experiment duration: %s | total ticks: %d", time.Since(startTime), pd.Ticks())
	// time.Sleep(10 * time.Second)

	f, err := os.OpenFile("snapshot.png", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	png.Encode(f, pd.TakeSnapshot())

	f.Close()
}
