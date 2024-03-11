package beer_test

import (
	"testing"

	"github.com/fmiskovic/go-consumer-pattern/internal/beer"
	mocks "github.com/fmiskovic/go-consumer-pattern/mocks/internal_/beer"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	//given
	barMock := mocks.NewBar(t)
	expected := "I want beer"
	barMock.EXPECT().IWantBeer().Return(expected)
	//when
	sut := beer.New(barMock)
	//then
	assert.Equal(t, expected, sut.Order())
}
