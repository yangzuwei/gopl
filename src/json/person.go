package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{"tom", 14}

	r, _ := json.Marshal(p)

	fmt.Printf("%s\n", r)
}
