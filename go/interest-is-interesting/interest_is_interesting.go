package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

	switch {
	case balance < 0:
		return 3.213
	case 0 <= balance && balance < 1000:
		return 0.5
	case 1000 <= balance && balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return (1 + float64(InterestRate(balance))/100) * balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	year := 0
	for ; balance < targetBalance; year += 1 {
		balance = AnnualBalanceUpdate(balance)
	}
	return year
}
