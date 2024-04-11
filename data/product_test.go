package data

import "testing"

func TestValidate(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 30.00,
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
