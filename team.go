package riley

import (
    "encoding/json"
)

// Team a type holding its characteristics.
type Team struct {
	Name      string
	Teammates []Teammate
}

// String returns the json representation of a dude.
func (t Team) String() string {
	raw, _ := json.Marshal(t)
	return string(raw)
}
