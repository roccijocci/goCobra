package todo

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
)

//Item Struct
type Item struct {
	Text string
}

//SaveItems Helperfunction
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

//ReadItems returns the struct Item
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	log.Println(err)
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}
