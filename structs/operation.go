package structs

// A Operation is a single GraphQL operation
type Operation struct {
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
}
