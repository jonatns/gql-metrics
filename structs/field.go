package structs

// An Argument is a single GrapQL Field argument
type Argument struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// A Field is a Operation field
type Field struct {
	Name      string     `json:"name"`
	Arguments []Argument `json:"arguments"`
	Fields    []Field    `json:"fields"`
}
