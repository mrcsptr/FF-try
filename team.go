package riley

import "encoding/json"

// Team a type holding its characteristics.
type Team struct {
	Name string
	Comp []Teammate
}
