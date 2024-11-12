package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type RecipesXML struct {
	Cakes []cake `xml:"cake"`
}

type RecipesJSON struct {
	Cakes []cake `json:"cake"`
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

func (i *ingredient) str() string {
	return fmt.Sprintf("%s: %s %s\n", i.Name, i.Count, i.Unit)
}

func (c *cake) str() string {
	s := fmt.Sprintf("%s\nStove Time: %s\nIngredients:\n", c.Name, c.StoveTime)

	for i, ingr := range c.Ingredients {
		s += fmt.Sprintf("%d) %s", i+1, ingr.str())
	}

	return s
}

func (r *RecipesXML) Read(file []byte) error {
	err := xml.Unmarshal(file, r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	return err
}

func (r *RecipesJSON) Read(file []byte) error {
	err := json.Unmarshal(file, r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	return err
}

func (r *RecipesXML) Print() {
	fmt.Printf("XML list of cakes:\n\n")
	for i, cake := range r.Cakes {
		fmt.Printf("%d. %s\n", i+1, cake.str())
	}
}

func (r *RecipesJSON) Print() {
	fmt.Printf("JSON list of cakes:\n\n")
	for i, cake := range r.Cakes {
		fmt.Printf("%d. %s\n", i+1, cake.str())
	}
}
