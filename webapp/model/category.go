package model

import (
	"fmt"
)

//Category struct
type Category struct {
	ID          int
	ImageURL    string
	Title       string
	Description string
}

var categories = []Category{
	Category{
		ID:       1,
		ImageURL: "lemon.png",
		Title:    "Juices and Mixes",
		Description: `Explore our wide assortment of juices and mixes expected by
							today's lemonade stand clientelle. Now featuring a full line of
							organic juices that are guaranteed to be obtained from trees that
							have never been treated with pesticides or artificial
							fertilizers.`,
	}, Category{
		ID:       2,
		ImageURL: "kiwi.png",
		Title:    "Cups, Straws, and Other Supplies",
		Description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
	}, Category{
		ID:       3,
		ImageURL: "pineapple.png",
		Title:    "Signs and Advertising",
		Description: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
	},
}

// GetCategories - returns all the categories
func GetCategories() []Category {
	return categories
}

// GetCategory - returns one requested category
func GetCategory(ID int) (*Category, error) {

	for _, c := range categories {
		if c.ID == ID {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("[category error]: category with ID: %v was not found", ID)
}
