package viewmodel

// StandLocator - struct
type StandLocator struct {
	Title  string
	Active string
}

// NewStandLocator - () - function
func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
	}
	return result
}
