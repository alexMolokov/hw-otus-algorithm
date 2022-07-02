package main

import (
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	sort "github.com/alexMolokov/hw-otus-algorithm/hw08_quick_sort"
)

func main() {
	pwd, _ := os.Getwd()
	for i := 2; i < 7; i++ {
		count := int(math.Pow(10, float64(i)))
		name := strconv.Itoa(count)
		p := pwd + filepath.FromSlash("/bin-data/"+name+"-values")
		err := sort.GenerateBinaryInt16(p, count)
		if err != nil {
			log.Fatalf("Can't generate binary file %v", err)
		}
	}
}
