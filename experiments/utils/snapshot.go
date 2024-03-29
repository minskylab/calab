package utils

import (
	"image/png"
	"os"

	"github.com/minskylab/calab/experiments/petridish"
)

func saveSnapshotToPNGImage(pd *petridish.PetriDish, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	return png.Encode(f, pd.TakeSnapshot())
}

func SaveSnapshotAsPNG(pd *petridish.PetriDish, filename string) error {
	return saveSnapshotToPNGImage(pd, filename)
}
