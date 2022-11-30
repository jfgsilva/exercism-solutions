package robotname

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name        string
	isNameTaken bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
	generator()
}

// maxNumberNames = possible combinations of 2 letters + 3 digits
const (
	maxNumberNames = 26 * 26 * 10 * 10 * 10
)

// cnt is increased each time a key is stored in db
var cnt = 0
var letters = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
var digits = strings.Split("0123456789", "")

// map pre created with max length to hold the robot names so we can check if they exist
var dbSlice = make([]Robot, maxNumberNames)

// anonymous func to generated the robot's name
func generator() {
	cnt2 := 0
	for _, l1 := range letters {
		for _, l2 := range letters {
			for _, d1 := range digits {
				for _, d2 := range digits {
					for _, d3 := range digits {
						name := l1 + l2 + d1 + d2 + d3
						dbSlice[cnt2] = Robot{name: name, isNameTaken: false}
						cnt2 += 1
					}
				}
			}
		}
	}
}

func (r *Robot) Name() (string, error) {
	// the name is assigned so no need to generate it: we just return the robot's name
	if r.name != "" {
		return r.name, nil
	}
	// infite loop
	for {
		if cnt >= maxNumberNames {
			// we ran through all possible values
			return "", errors.New("max Number of different names reached")
		}
		randRobot := rand.Intn(maxNumberNames)
		if !dbSlice[randRobot].isNameTaken {
			dbSlice[randRobot].isNameTaken = true
			r.name = dbSlice[randRobot].name
			cnt += 1
			return r.name, nil
		} else {
			continue
		}
	}

}

func (r *Robot) Reset() {
	r.name = ""
}
