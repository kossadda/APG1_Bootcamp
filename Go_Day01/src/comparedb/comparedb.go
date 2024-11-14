// Package comparedb provides functions for comparing recipe databases.
package comparedb

import (
	"fmt"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

// Compare compares two recipe databases and prints the differences.
func Compare(old recipes.Recipes, new recipes.Recipes) {
	fmt.Print(compareCakes(old, new))
	fmt.Print(compareTimes(old, new))
	fmt.Print(compareIngredients(old, new))
	fmt.Print(compareUnits(old, new))
}

// compareCakes compares the cakes in two recipe databases.
func compareCakes(old recipes.Recipes, new recipes.Recipes) (comp string) {
	oldCakes := make(map[string]bool)
	newCakes := make(map[string]bool)

	for _, cake := range old.Cakes {
		oldCakes[cake.Name] = true
	}

	for _, cake := range new.Cakes {
		newCakes[cake.Name] = true
	}

	for cakeName := range newCakes {
		if !oldCakes[cakeName] {
			comp += fmt.Sprintf("ADDED cake \"%s\"\n", cakeName)
		}
	}

	for cakeName := range oldCakes {
		if !newCakes[cakeName] {
			comp += fmt.Sprintf("REMOVED cake \"%s\"\n", cakeName)
		}
	}

	return comp
}

// compareTimes compares the cooking times of cakes in two recipe databases.
func compareTimes(old recipes.Recipes, new recipes.Recipes) (comp string) {
	oldCakes := make(map[string]string)
	newCakes := make(map[string]string)

	for _, cake := range old.Cakes {
		oldCakes[cake.Name] = cake.StoveTime
	}

	for _, cake := range new.Cakes {
		newCakes[cake.Name] = cake.StoveTime
	}

	for cakeName, newTime := range newCakes {
		if oldTime, exists := oldCakes[cakeName]; exists && oldTime != newTime {
			comp += fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cakeName, newTime, oldTime)
		}
	}

	return comp
}

// compareIngredients compares the ingredients of cakes in two recipe databases.
func compareIngredients(old recipes.Recipes, new recipes.Recipes) (comp string) {
	oldCakes := make(map[string]recipes.Cake)
	newCakes := make(map[string]recipes.Cake)

	for _, cake := range old.Cakes {
		oldCakes[cake.Name] = cake
	}

	for _, cake := range new.Cakes {
		newCakes[cake.Name] = cake
	}

	for cakeName, newCake := range newCakes {
		if oldCake, exists := oldCakes[cakeName]; exists {
			comp += addedIngredients(oldCake, newCake)
			comp += removedIngredients(oldCake, newCake)
		}
	}

	return comp
}

// compareUnits compares the units of ingredients in two recipe databases.
func compareUnits(old recipes.Recipes, new recipes.Recipes) (comp string) {
	oldCakes := make(map[string]recipes.Cake)
	newCakes := make(map[string]recipes.Cake)

	for _, cake := range old.Cakes {
		oldCakes[cake.Name] = cake
	}

	for _, cake := range new.Cakes {
		newCakes[cake.Name] = cake
	}

	for cakeName, newCake := range newCakes {
		if oldCake, exists := oldCakes[cakeName]; exists {
			comp += changedUnit(oldCake, newCake)
		}
	}

	return comp
}

// changedUnit compares the units of ingredients in two cakes.
func changedUnit(old recipes.Cake, new recipes.Cake) (comp string) {
	for _, newIngr := range new.Ingredients {
		for _, oldIngr := range old.Ingredients {
			if newIngr.Name == oldIngr.Name {
				if newIngr.Unit != oldIngr.Unit {
					if newIngr.Unit == "" {
						comp += fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngr.Unit, oldIngr.Name, old.Name)
					} else {
						comp += fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngr.Name, old.Name, newIngr.Unit, oldIngr.Unit)
					}
				} else if newIngr.Count != oldIngr.Count {
					comp += fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngr.Name, old.Name, newIngr.Count, oldIngr.Count)
				}
			}
		}
	}

	return comp
}

// addedIngredients finds and returns the added ingredients in the new cake.
func addedIngredients(old recipes.Cake, new recipes.Cake) (comp string) {
	addedIngr := 0

	for _, newIngr := range new.Ingredients {
		for _, oldIngr := range old.Ingredients {
			if newIngr.Name != oldIngr.Name {
				addedIngr++
			}
		}

		if addedIngr == len(new.Ingredients) {
			comp += fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngr.Name, old.Name)
		}

		addedIngr = 0
	}

	return comp
}

// removedIngredients finds and returns the removed ingredients in the old cake.
func removedIngredients(old recipes.Cake, new recipes.Cake) (comp string) {
	removeIngr := 0

	for _, oldIngr := range old.Ingredients {
		for _, newIngr := range new.Ingredients {
			if oldIngr.Name != newIngr.Name {
				removeIngr++
			}
		}

		if removeIngr == len(old.Ingredients) {
			comp += fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngr.Name, old.Name)
		}

		removeIngr = 0
	}

	return comp
}
