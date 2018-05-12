package gomarvel

import (
	"net/http"

	"github.com/adrianjjohnson/gomarvel/model"
)

const (
	baseURI     = "http://gateway.marvel.com"
	routePrefix = "/v1/public"
	charactsers = "/characters"
	comics      = "/comics"
)

// Service defines the type of resources this package will implement.
type Service interface {
	CharacterService
	ComicService
}

// CharacterService is responsible for fetching character data.
type CharacterService interface {
	GetAllCharacters() (model.CharacterResult, error)
	GetCharacterByID(id int64) (model.CharacterResult, error)
}

// ComicService is responsible for fetching comic data.
type ComicService interface {
	GetComicByCharacterID(id int64) (model.ComicResult, error)
}

// Client contains resources necessary to contact the Marvel API.
type Client struct {
	publicKey  string
	privateKey string
	client     *http.Client
}

// New initializes a new Service
func New(publicKey, privateKey string) Service {
	return &Client{
		publicKey,
		privateKey,
		&http.Client{},
	}
}
