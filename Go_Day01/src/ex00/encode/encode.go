package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Recipes struct {
	Cakes []cake `xml:"cake" json:"cake"`
}

type cake struct {
	Name        string       `xml:"name" json:"name"`
	StoveTime   string       `xml:"stovetime" json:"time"`
	Ingredients []ingredient `xml:"ingredients>item" json:"ingredients"`
}

type ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit"`
}

type RecipesXML struct{}

type RecipesJSON struct{}

func (r *RecipesXML) Read(file []byte) (Recipes, error) {
	var recipes Recipes
	err := xml.Unmarshal(file, &recipes)
	if err != nil {
		return Recipes{}, err
	}
	
	return recipes, nil
}

func (r *RecipesJSON) Read(file []byte) (Recipes, error) {
	var recipes Recipes
	err := json.Unmarshal(file, &recipes)
	if err != nil {
		return Recipes{}, err
	}

	return recipes, nil
}

func (r *RecipesXML) Print(recipes Recipes) {
	jsonData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func (r *RecipesJSON) Print(recipes Recipes) {
	xmlData, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
