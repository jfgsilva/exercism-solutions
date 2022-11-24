package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`(\W|^)^\[ERR]|^\[TRC]|^\[DBG]|^\[INF]|^\[WRN]|^\[FTL](\W)`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*`)
	counts := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counts += 1
		}
	}
	return counts
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +(\w+)`)
	res := make([]string, len(lines))
	copy(res, lines)
	for i, line := range lines {
		name := re.FindStringSubmatch(line)
		if len(name) == 2 {
			res[i] = fmt.Sprintf("[USR] %s %s", name[1], line)
		}
	}
	return res
}
