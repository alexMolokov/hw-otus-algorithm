package sort

import (
	"encoding/binary"
	"math/rand"
	"os"
	"time"
)

const (
	MaxRecordInt16 = 65536 / 2
	MaxValueInt16  = 65535
)

func GenerateBinaryInt16(fileName string, count int) error {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0666) // nolint
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	var data []int16
	if count > MaxRecordInt16 {
		data = make([]int16, MaxRecordInt16)
	} else {
		data = make([]int16, count)
	}

	rand.Seed(time.Now().UnixNano())

	var i, current int
	for i < count {
		for current < len(data) {
			data[current] = int16(rand.Intn(MaxValueInt16)) //nolint
			current++
			i++
		}
		err := binary.Write(file, binary.BigEndian, data[0:current])
		if err != nil {
			return err
		}
		current = 0
	}

	return nil
}
