package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	j01, err01 := json.Marshal(u1)
	j02, err02 := json.Marshal(u2)
	j03, err03 := json.Marshal(u3)


	if err01 != nil{
		fmt.Printf("%s\n",err01)
	} else {
		fmt.Printf("%s\n",j01)
	}
	if err02 != nil{
		fmt.Printf("%s\n",err02)
	} else {
		fmt.Printf("%s\n",j02)
	}
	if err03 != nil{
		fmt.Printf("%s\n",err03)
	} else {
		fmt.Printf("%s\n",j03)
	}

}