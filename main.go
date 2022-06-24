package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// identical expression in logical comparitor
	// empty branch
	if 1 == 1 {

	}
	//duplicate code
	if 1 == 1 {

	}

	//error not handled
	json, _ := json.Marshal("marshal_me")
	//Output not used
	fmt.Sprint(json)

	//loop condition never changes
	for i := 0; i < 1; i++ {
		panic("I don't like loops")
	}
}
