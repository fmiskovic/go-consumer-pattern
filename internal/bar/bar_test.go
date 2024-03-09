package bar_test

import (
	"testing"

	"github.com/fmiskovic/go-consumer-pattern/internal/bar"
	"github.com/stretchr/testify/assert"
)

func TestIWantBeer(t *testing.T) {
	bar := bar.Bar{}
	assert.Equal(t, "I want beer", bar.IWantBeer())
}

func TestIWantWine(t *testing.T) {
	bar := bar.Bar{}
	assert.Equal(t, "I want wine", bar.IWantWine())
}

func TestIWantWhisky(t *testing.T) {
	bar := bar.Bar{}
	assert.Equal(t, "I want whisky", bar.IWantWhisky())
}
