package main

import (
	"errors"

	"fmt"
)

func main() {

	/*slice1 := make([]int, 100)
	  for i, _ := range slice1 {
	      slice1[i] = rand.Intn(50)

	  }

	  slice1 := []int{4, 5, 6, 7, 1, 2, 3, 4, 1}
	  mymap := make(map[int]uint16)
	  for _, v := range slice1 {

	      if _, ok := mymap[v]; ok {

	          mymap[v] = mymap[v] + 1

	      } else {

	          mymap[v] = 1

	      }

	  }

	  fmt.Println(mymap)*/

	my := map[string]int{"megha": 1, "pratheeksha": 2}

	Key := "megha"

	Bool, err := Delete(my, Key)

	fmt.Println(Bool, err)

	//key1 := "hii"

	//	bool1, err1 := Update(my, key1)

	fmt.Println()

}

func Delete(my map[string]int, key string) (bool, error) {

	if my == nil {

		return false, errors.New("map is nill")

	}

	if _, ok := my[key]; ok {

		delete(my, key)

		return true, nil

	} else {

		return false, nil

	}

}

func Update(my map[string]int, key string) (bool, error) {

	if my == nil {

		return false, errors.New("map is nill")

	}

	if _, ok := my[key]; ok {

		delete(my, key)

		return true, nil

	} else {

		return false, errors.New("key doesntexist")

	}

}
