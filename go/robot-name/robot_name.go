package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const maxNumberNames = 26 * 26 * 10 * 10 * 10

var Cnt = 0

var DB = make(map[string]int, maxNumberNames)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	generator := func() string {
		// generates two letters
		ltrs := make([]rune, 2)
		for i := range ltrs {
			ltrs[i] = letters[rand.Intn(len(letters))]
		}
		tripe_digit := fmt.Sprintf("%03d", rand.Intn(1000))
		return string(ltrs) + tripe_digit
	}
	for {
		key := generator()
		if Cnt > maxNumberNames {
			return "", errors.New("max Number of different names reached")
		}
		if _, ok := DB[key]; !ok {
			DB[key] = 1
			Cnt += 1
			r.name = key
			return key, nil
		} else {
			continue
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
}
