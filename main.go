// main.go

package main

import (
	"fmt"
	"dicitionnaires/dictionary"
)

func main() {
	dictionary := dictionary.NewDictionary()

	
	dictionary.Add("go", "langage de programmation")
	dictionary.Add("python", " script")
	dictionary.Add("java", " orienté objet")

	definition, found := dictionary.Get("go")
	if found {
		fmt.Printf("La définition de 'go' est : %s\n", definition)
	} else {
		fmt.Println("Mot non trouvé dans le dictionnaire.")
	}

	dictionary.Remove("python")

	dictionary.List()
}
