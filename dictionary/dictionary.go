

package dictionary

import (
	"fmt"
	"sort"
)


type Dictionary struct {
	entries map[string]string
}


func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}


func (d *Dictionary) Add(word, definition string) {
	d.entries[word] = definition
}

func (d *Dictionary) Get(word string) (string, bool) {
	definition, found := d.entries[word]
	return definition, found
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}


func (d *Dictionary) List() {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	sort.Strings(words)

	fmt.Println("Liste des mots et de leurs définitions :")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, d.entries[word])
	}
}
