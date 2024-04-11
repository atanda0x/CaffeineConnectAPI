package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 30.00,
	}

	err := p.Validate()
	require.NoError(t, err)
}
