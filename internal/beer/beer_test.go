package beer_test

import (
	"testing"

	"github.com/fmiskovic/go-consumer-pattern/internal/beer"
	mocks "github.com/fmiskovic/go-consumer-pattern/mocks/internal_/beer"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	barMock := mocks.NewBar(t)
	barMock.EXPECT().IWantBeer().Return("I want beer")

	sut := beer.New(barMock)
	assert.Equal(t, "I want beer", sut.Order())
}
