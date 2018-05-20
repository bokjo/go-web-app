package viewmodel

// Base type
type Base struct {
	Title string
}

// NewBase constructor function
func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
