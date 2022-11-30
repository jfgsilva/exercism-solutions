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

// maxNumberNames = possible combinations of 2 letters + 3 digits
const maxNumberNames = 26 * 26 * 10 * 10 * 10

// Cnt is increased each time a key is stored in DB
var Cnt = 0

// map pre created with max length to hold the robot names so we can check if they exist
var DB = make(map[string]int, maxNumberNames)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (r *Robot) Name() (string, error) {
	// the name is assigned so no need to generate it: we just return the robot's name
	if r.name != "" {
		return r.name, nil
	}
	// anonymous func to generated the robot's name
	generator := func() string {
		// generates two letters
		ltrs := make([]rune, 2)
		for i := range ltrs {
			ltrs[i] = letters[rand.Intn(len(letters))]
		}
		///generates three digits
		tripe_digit := fmt.Sprintf("%03d", rand.Intn(1000))
		return string(ltrs) + tripe_digit
	}
	// infite loop
	for {
		key := generator()
		if Cnt > maxNumberNames {
			// we ran through all possible values
			return "", errors.New("max Number of different names reached")
		}
		// we check if
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
