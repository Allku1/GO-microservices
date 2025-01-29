package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "allkui",
		Price: 1.00,
		SKU:   "avsasd-afafas-dsgsdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
