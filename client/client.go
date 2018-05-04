package client

import (
	"net/http"

	"github.com/adrianjjohnson/gomarvel"
)

const (
	baseURI     = "http://gateway.marvel.com"
	routePrefix = "/v1/public"
	charactsers = "/characters"
)

type Service interface {
	CharacterService
}

// CharacterService is responsible for fetching charactes.
type CharacterService interface {
	GetAllCharacters() (gomarvel.Result, error)
}

type Client struct {
	publicKey  string
	privateKey string
	client     *http.Client
}

func New(publicKey, privateKey string) Service {
	return &Client{
		publicKey,
		privateKey,
		&http.Client{},
	}
}
