package main

import (
	"errors"
	"fmt"
)

func main() {
	myMap := map[string]int{
		"megha":       1,
		"pratheeksha": 2,
	}

	key := "megha"
	newValue := 3
	boolVal, err := Update(myMap, key, newValue)
	fmt.Println(boolVal, err)
	fmt.Println(myMap)
}

func Delete(myMap map[string]int, key string) (bool, error) {
	if myMap == nil {
		return false, errors.New("map is nil")
	}

	if _, ok := myMap[key]; ok {
		delete(myMap, key)
		return true, nil
	} else {
		return false, nil
	}
}

func Update(myMap map[string]int, key string, newValue int) (bool, error) {
	if myMap == nil {
		return false, errors.New("map is nil")
	}

	if _, ok := myMap[key]; ok {
		myMap[key] = newValue
		return true, nil
	} else {
		return false, errors.New("key doesn't exist")
	}
}
