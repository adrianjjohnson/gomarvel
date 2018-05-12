package gomarvel

import (
	"net/http"
	"os"
	"testing"
)

func TestGetAllCharacters(t *testing.T) {
	svc := newTestClient()
	c, err := svc.GetAllCharacters()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Results: %v", c.Data.Results)
}

func TestGetCharacterByID(t *testing.T) {
	result := []struct {
		id   int64
		want string
	}{
		{1011334, "3-D Man"},
		{1017100, "A-Bomb (HAS)"},
	}

	svc := newTestClient()

	for _, v := range result {
		c, err := svc.GetCharacterByID(v.id)
		if err != nil {
			t.Error(err)
		}

		for _, i := range c.Data.Results {
			if i.Name != v.want {
				t.Errorf("got: %s, want: %s", i.Name, v.want)
			}
			t.Log(i.Name)
		}
	}
}
func newTestClient() *Client {
	return &Client{
		client:     http.DefaultClient,
		publicKey:  os.Getenv("PUBLIC_KEY"),
		privateKey: os.Getenv("PRIVATE_KEY"),
	}
}
