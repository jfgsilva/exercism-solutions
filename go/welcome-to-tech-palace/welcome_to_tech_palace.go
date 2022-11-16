package techpalace
import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	s := strings.Repeat("*", numStarsPerLine) + "\n"
    s += welcomeMsg + "\n"
    s += strings.Repeat("*", numStarsPerLine)
    return s
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	ch := []string{"\n", "\t", "*"}
    for _, c := range ch {
        oldMsg = strings.ReplaceAll(oldMsg, c, "")
    }
return strings.TrimSpace(oldMsg)
    }
