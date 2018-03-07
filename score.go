package riley

import (
	"fmt"
	"strings"
)

// Definition of the values for a win (W), tie (T), and a loss (L).
const (
	W = 3
	T = 2
	L = 1
)

// ParseScore assigns values defined as constants and returns an error if the character is missing or not what was expected.
func ParseScore(raw string) (int, error) {
	switch {
	case strings.EqualFold(raw, "W"):
		return W, nil
	case strings.EqualFold(raw, "T"):
		return T, nil
	case strings.EqualFold(raw, "L"):
		return L, nil
	default:
		return 0, fmt.Errorf("invalid score '%s'", raw)
	}
}
