package main

import (
	"fmt"
	"os"

	"github.com/adrianjjohnson/gomarvel"
)

func main() {
	publicKey := os.Getenv("PUBLIC_KEY")
	privateKey := os.Getenv("PRIVATE_KEY")

	svc := gomarvel.New(publicKey, privateKey)

	character, err := svc.GetAllCharacters()
	if err != nil {
		panic(err)
	}

	for _, i := range character.Data.Results {
		fmt.Println(i.Name)
	}

	c, err := svc.GetCharacterByID(1011334)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Data.Results)
}
