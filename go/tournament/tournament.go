package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
	points        int
}

func (t *team) win() {
	t.matchesPlayed += 1
	t.wins += 1
	t.points += 3
}

func (t *team) loss() {
	t.matchesPlayed += 1
	t.losses += 1
}

func (t *team) draw() {
	t.matchesPlayed += 1
	t.points += 1
	t.draws += 1
}

func (t team) String() string {
	s := fmt.Sprintf("%-30s |%3v |%3v |%3v |%3v |%3v", t.name, t.matchesPlayed, t.wins, t.draws, t.losses, t.points)
	// fmt.Println("size of s", len(s))
	return s
}

type scoreboard map[string]*team

func (s scoreboard) team(name string) *team {
	// we check if key exists
	t, ok := s[name]
	// if not we create the pointer, and initialize the name
	if !ok {
		//t.name = name
		t = &team{name: name}
		s[name] = t
	}
	return t
}

func (s scoreboard) String() string {
	var slcTeams []team
	for _, val := range s {
		slcTeams = append(slcTeams, *val)
	}
	sort.Slice(slcTeams, func(i, j int) bool {
		if slcTeams[i].points > slcTeams[j].points {
			return true
		} else if slcTeams[i].points == slcTeams[j].points {
			return slcTeams[i].name < slcTeams[j].name
		}
		return false
	})
	header := fmt.Sprintf("%-30s |%3v |%3v |%3v |%3v |%3v", "Team", "MP", "W", "D", "L", "P")
	var str = []string{header}
	for _, line := range slcTeams {
		str = append(str, fmt.Sprint(line))
	}
	// str = append(str, "\n")
	return strings.Join(str, "\n") + "\n"
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	sb := make(scoreboard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" || line[:1] == "#" {
			// fmt.Printf("### excluded: %s\n", line)
			continue
		}
		// fmt.Printf("%s", line)
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) != 3 {
			return errors.New("wrong size for record")
		}

		t1, t2, result := fields[0], fields[1], fields[2]
		// if the result is not expected return error
		// if !(result == "draw" || result == "win" || result == "loss") {
		// 	return errors.New("result is not correct")
		// }
		// fmt.Printf("t1: %s, t2: %s, res: %s\n", t1, t2, result)
		switch result {
		case "win":
			sb.team(t1).win()
			sb.team(t2).loss()
		case "loss":
			sb.team(t1).loss()
			sb.team(t2).win()
		case "draw":
			sb.team(t1).draw()
			sb.team(t2).draw()
		default:
			return errors.New("result is not correct")
		}
	}
	// fmt.Println(sb)
	writer.Write([]byte(sb.String()))
	return nil
}
