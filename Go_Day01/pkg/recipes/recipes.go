// Package recipes provides types for representing recipe data.
package recipes

// Recipes represents a collection of cake recipes.
type Recipes struct {
	Cakes []Cake `xml:"cake" json:"cake"`
}

// Cake represents a single cake recipe.
type Cake struct {
	Name        string       `xml:"name" json:"name"`
	StoveTime   string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

// Ingredient represents an ingredient in a cake recipe.
type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit"`
}
