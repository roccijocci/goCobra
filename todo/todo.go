package todo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"io/ioutil"
	"log"
)

//Item Struct
type Item struct {
	Text     string
	Priority int
	position int
}

//SetPriority casee
func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

//PrettyP to make the output pretty???
func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
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
