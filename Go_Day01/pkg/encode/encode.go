// Package encode provides functions for encoding and decoding recipe data.
package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/recipes"
)

// RecipesXML is a struct for reading XML recipe data.
type RecipesXML struct {
	*recipes.Recipes
}

// RecipesJSON is a struct for reading JSON recipe data.
type RecipesJSON struct {
	*recipes.Recipes
}

// Read reads XML recipe data from the given file.
func (r *RecipesXML) DBRead(file []byte) error {
	r.Recipes = &recipes.Recipes{}
	return xml.Unmarshal(file, r.Recipes)
}

// Read reads JSON recipe data from the given file.
func (r *RecipesJSON) DBRead(file []byte) error {
	r.Recipes = &recipes.Recipes{}
	return json.Unmarshal(file, r.Recipes)
}

// Print prints the recipe data in JSON format.
func (r *RecipesXML) String() string {
	jsonData, err := json.MarshalIndent(r.Recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return ""
	}

	return string(jsonData)
}

// Print prints the recipe data in XML format.
func (r *RecipesJSON) String() string {
	xmlData, err := xml.MarshalIndent(r.Recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return ""
	}

	return string(xmlData)
}

func (r *RecipesXML) Recipe() recipes.Recipes {
	return *r.Recipes
}

func (r *RecipesJSON) Recipe() recipes.Recipes {
	return *r.Recipes
}
