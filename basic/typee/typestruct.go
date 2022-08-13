package typee

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func MainTypeStruct() {
	rec1 := Rectangle{Width: 20, Height: 40}
	j, err := json.Marshal(&rec1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))

	var rec2 Rectangle
	if err := json.Unmarshal([]byte(j), &rec2); err != nil {
		log.Panic(err)
	}
	fmt.Printf("%#v", rec2)
}
