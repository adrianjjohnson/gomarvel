# gomarvel

## Description
Package connects to the Marvel official API

## Getting started
Using this package requres a valid Marvel API public and private key

[Marvel API Documentation](https://developer.marvel.com/)

## Instalation
```go get -u github.com/adrianjjohnson/gomarvel```

## Example
```go
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
}
```