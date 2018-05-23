package viewmodel

// Base type
type Base struct {
	Active string
	Title  string
}

// NewBase constructor function
func NewBase() Base {
	return Base{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
}
