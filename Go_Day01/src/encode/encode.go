package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

type RecipesXML struct{}

type RecipesJSON struct{}

func (r *RecipesXML) Read(file []byte) (recipes.Recipes, error) {
	var recipes recipes.Recipes
	err := xml.Unmarshal(file, &recipes)

	return recipes, err
}

func (r *RecipesJSON) Read(file []byte) (recipes.Recipes, error) {
	var recipes recipes.Recipes
	err := json.Unmarshal(file, &recipes)

	return recipes, err
}

func (r *RecipesXML) Print(recipes recipes.Recipes) {
	jsonData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func (r *RecipesJSON) Print(recipes recipes.Recipes) {
	xmlData, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
