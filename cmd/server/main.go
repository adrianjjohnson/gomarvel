package main

import (
	"fmt"
	"os"

	gmarvel "github.com/adrianjjohnson/gomarvel/client"
)

func main() {
	publicKey := os.Getenv("PUBLIC_KEY")
	privateKey := os.Getenv("PRIVATE_KEY")
	svc := gmarvel.New(publicKey, privateKey)

	character, err := svc.GetAllCharacters()
	if err != nil {
		panic(err)
	}

	for _, i := range character.Data.Results {
		fmt.Println(i.Name)
	}
}
