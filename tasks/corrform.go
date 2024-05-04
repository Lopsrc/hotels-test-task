package tasks

import "fmt"

// компьютер  компьютера    компьютеров
// 1 21 31... 2-4,22-24... 5-20, 25-30...

// GetStringInCorrectForm returns a string representation of a number in the correct form for Russian language.
func GetStringInCorrectForm(n int) string{
	m := n%10
	switch {
	case m >= 5 && m <= 9 || m == 0 || n>=11 && n<=14:
		return fmt.Sprintf("%d компьютеров", n)
	case m >=2 && m <= 4:
		return fmt.Sprintf("%d компьютера", n)
	case m == 1:
		return fmt.Sprintf("%d компьютер", n)
	}
	return ""
}