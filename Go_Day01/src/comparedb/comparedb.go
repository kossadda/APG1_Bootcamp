package comparedb

import (
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

func CompareRecipes(old recipes.Recipes, new recipes.Recipes) {
	fmt.Print(compareCakes(old, new))
	fmt.Print(compareTimes(old, new))
	fmt.Print(compareIngredients(old, new))
	fmt.Print(compareUnits(old, new))
}

func compareCakes(old recipes.Recipes, new recipes.Recipes) (comp string) {
	old_сakes := make(map[string]bool)
	new_сakes := make(map[string]bool)

	for _, cake := range old.Cakes {
		old_сakes[cake.Name] = true
	}

	for _, cake := range new.Cakes {
		new_сakes[cake.Name] = true
	}

	for cake_name := range new_сakes {
		if !old_сakes[cake_name] {
			comp += fmt.Sprintf("ADDED cake \"%s\"\n", cake_name)
		}
	}

	for cake_name := range old_сakes {
		if !new_сakes[cake_name] {
			comp += fmt.Sprintf("REMOVED cake \"%s\"\n", cake_name)
		}
	}

	return comp
}

func compareTimes(old recipes.Recipes, new recipes.Recipes) (comp string) {
	old_сakes := make(map[string]string)
	new_сakes := make(map[string]string)

	for _, cake := range old.Cakes {
		old_сakes[cake.Name] = cake.StoveTime
	}

	for _, cake := range new.Cakes {
		new_сakes[cake.Name] = cake.StoveTime
	}

	for cake_name, new_time := range new_сakes {
		if old_time, exists := old_сakes[cake_name]; exists && old_time != new_time {
			comp += fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake_name, new_time, old_time)
		}
	}

	return comp
}

func compareIngredients(old recipes.Recipes, new recipes.Recipes) (comp string) {
	old_сakes := make(map[string]recipes.Cake)
	new_сakes := make(map[string]recipes.Cake)

	for _, cake := range old.Cakes {
		old_сakes[cake.Name] = cake
	}

	for _, cake := range new.Cakes {
		new_сakes[cake.Name] = cake
	}

	for cake_name, new_cake := range new_сakes {
		if old_cake, exists := old_сakes[cake_name]; exists {
			comp += addedIngredients(old_cake, new_cake)
			comp += removedIngredients(old_cake, new_cake)
		}
	}

	return comp
}

func compareUnits(old recipes.Recipes, new recipes.Recipes) (comp string) {
	old_сakes := make(map[string]recipes.Cake)
	new_сakes := make(map[string]recipes.Cake)

	for _, cake := range old.Cakes {
		old_сakes[cake.Name] = cake
	}

	for _, cake := range new.Cakes {
		new_сakes[cake.Name] = cake
	}

	for cake_name, new_cake := range new_сakes {
		if old_cake, exists := old_сakes[cake_name]; exists {
			comp += changedUnit(old_cake, new_cake)
		}
	}

	return comp
}

func changedUnit(old recipes.Cake, new recipes.Cake) (comp string) {
	for _, new_ingr := range new.Ingredients {
		for _, old_ingr := range old.Ingredients {
			if new_ingr.Name == old_ingr.Name {
				if new_ingr.Unit != old_ingr.Unit {
					if new_ingr.Unit == "" {
						comp += fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", old_ingr.Unit, old_ingr.Name, old.Name)
					} else {
						comp += fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", new_ingr.Name, old.Name, new_ingr.Unit, old_ingr.Unit)
					}
				} else if new_ingr.Count != old_ingr.Count {
					comp += fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", new_ingr.Name, old.Name, new_ingr.Count, old_ingr.Count)
				}
			}
		}
	}

	return comp
}

func addedIngredients(old recipes.Cake, new recipes.Cake) (comp string) {
	added_ingr := 0

	for _, new_ingr := range new.Ingredients {
		for _, old_ingr := range old.Ingredients {
			if new_ingr.Name != old_ingr.Name {
				added_ingr++
			}
		}

		if added_ingr == len(new.Ingredients) {
			comp += fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n", new_ingr.Name, old.Name)
		}

		added_ingr = 0
	}

	return comp
}

func removedIngredients(old recipes.Cake, new recipes.Cake) (comp string) {
	remove_ingr := 0

	for _, old_ingr := range old.Ingredients {
		for _, new_ingr := range new.Ingredients {
			if old_ingr.Name != new_ingr.Name {
				remove_ingr++
			}
		}

		if remove_ingr == len(old.Ingredients) {
			comp += fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"\n", old_ingr.Name, old.Name)
		}

		remove_ingr = 0
	}

	return comp
}
