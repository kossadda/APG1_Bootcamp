// Package encode provides functions for encoding and decoding recipe data.
package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

// RecipesXML is a struct for reading XML recipe data.
type RecipesXML struct{}

// RecipesJSON is a struct for reading JSON recipe data.
type RecipesJSON struct{}

// Read reads XML recipe data from the given file.
func (r *RecipesXML) Read(file []byte) (recipes.Recipes, error) {
	var recipes recipes.Recipes
	err := xml.Unmarshal(file, &recipes)

	return recipes, err
}

// Read reads JSON recipe data from the given file.
func (r *RecipesJSON) Read(file []byte) (recipes.Recipes, error) {
	var recipes recipes.Recipes
	err := json.Unmarshal(file, &recipes)

	return recipes, err
}

// Print prints the recipe data in JSON format.
func (r *RecipesXML) Print(recipes recipes.Recipes) {
	jsonData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

// Print prints the recipe data in XML format.
func (r *RecipesJSON) Print(recipes recipes.Recipes) {
	xmlData, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
