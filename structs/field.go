package structs

// A Field is a Operation field
type Field struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}
