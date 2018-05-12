package gomarvel

import "testing"

func TestGetAllComics(t *testing.T) {
	svc := newTestClient()
	c, err := svc.GetAllComics()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Results: %v", c.Data.Results)
}

func TestGetComicByCharacterID(t *testing.T) {
	result := []struct {
		id   int64
		want string
	}{
		{1011334, "Avengers: The Initiative (2007) #19"},
	}

	svc := newTestClient()

	for _, v := range result {
		c, err := svc.GetComicByCharacterID(v.id)
		if err != nil {
			t.Error(err)
		}
		t.Log(c.Data.Results[0])
	}
}
