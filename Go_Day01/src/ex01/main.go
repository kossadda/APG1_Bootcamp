package main

import (
	"flag"
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/readdb"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
	"os"
)

func main() {
	var reader readdb.DBReader
	var old recipes.Recipes
	var new recipes.Recipes
	old_path := flag.String("old", "", "Original database (xml or json)")
	new_path := flag.String("new", "", "Stolen database (xml or json)")
	flag.Parse()

	old_file, err := readdb.DefineFile(&reader, old_path)

	if err == nil {
		old, err = reader.Read(old_file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	new_file, err := readdb.DefineFile(&reader, new_path)

	if err == nil {
		new, err = reader.Read(new_file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	compareRecipes(old, new)
}

func compareRecipes(old recipes.Recipes, new recipes.Recipes) {
	compareCakes(old, new)
	compareTimes(old, new)
	compareIngredients(old, new)
}

func compareCakes(old recipes.Recipes, new recipes.Recipes) {
	added_cakes := 0
	remove_cakes := 0
	
	for _, new_cake := range new.Cakes {
		for _, old_cake := range old.Cakes {
			if new_cake.Name != old_cake.Name {
				added_cakes++
			}
		}
		
		if added_cakes == len(old.Cakes) {
			fmt.Printf("ADDED cake \"%s\"\n", new_cake.Name)
		}
		
		added_cakes = 0
	}
	
	for _, old_cake := range old.Cakes {
		for _, new_cake := range new.Cakes {
			if old_cake.Name != new_cake.Name {
				remove_cakes++
			}
		}
		
		if remove_cakes == len(old.Cakes) {
			fmt.Printf("REMOVED cake \"%s\"\n", old_cake.Name)
		}
		
		remove_cakes = 0
	}
}

func compareTimes(old recipes.Recipes, new recipes.Recipes) {
	for _, old_cake := range old.Cakes {
		for _, new_cake := range new.Cakes {
			if old_cake.Name == new_cake.Name {
				if old_cake.StoveTime != new_cake.StoveTime {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", old_cake.Name, new_cake.StoveTime, old_cake.StoveTime)
				}
			}
		}
	}
}

func compareIngredients(old recipes.Recipes, new recipes.Recipes) {
	for _, new_cake := range new.Cakes {
		for _, old_cake := range old.Cakes {
			if new_cake.Name == old_cake.Name {
				addedIngredients(old_cake, new_cake)
			}
		}
	}
	
	for _, old_cake := range old.Cakes {
		for _, new_cake := range new.Cakes {
			if old_cake.Name == new_cake.Name {
				removedIngredients(old_cake, new_cake)
			}
		}
	}
}

func addedIngredients(old recipes.Cake, new recipes.Cake) {
	added_ingr := 0

	for _, new_ingr := range new.Ingredients {
		for _, old_ingr := range old.Ingredients {
			if new_ingr.Name != old_ingr.Name {
				added_ingr++
			}
		}

		if added_ingr == len(new.Ingredients) {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", new_ingr.Name, old.Name)
		}

		added_ingr = 0
	}
}

func removedIngredients(old recipes.Cake, new recipes.Cake) {
	remove_ingr := 0

	for _, old_ingr := range old.Ingredients {
		for _, new_ingr := range new.Ingredients {
			if old_ingr.Name != new_ingr.Name {
				remove_ingr++
			}
		}

		if remove_ingr == len(old.Ingredients) {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", old_ingr.Name, old.Name)
		}

		remove_ingr = 0
	}
}

// func compareUnits(old recipes.Recipes, new recipes.Recipes) {
// 	for _, new_cake := range new.Cakes {
// 		for _, old_cake := range old.Cakes {
// 			if new_cake.Name == old_cake.Name {
// 				addedIngredients(old_cake, new_cake)
// 			}
// 		}
// 	}
	
// 	for _, old_cake := range old.Cakes {
// 		for _, new_cake := range new.Cakes {
// 			if old_cake.Name == new_cake.Name {
// 				removedIngredients(old_cake, new_cake)
// 			}
// 		}
// 	}
// }

// func changedUnit(old recipes.Ingredient, new recipes.Ingredient) {
// 	for _, new_ingr := range new.Ingredients {
// 		for _, old_ingr := range old.Ingredients {
// 			if new_ingr.Name == old_ingr.Name {
				
// 			}
// 		}

// 		if added_ingr == len(new.Ingredients) {
// 			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", new_ingr.Name, old.Name)
// 		}

// 		added_ingr = 0
// 	}
// }