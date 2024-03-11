package wine_test

import (
	"testing"

	"github.com/fmiskovic/go-consumer-pattern/internal/wine"
	mocks "github.com/fmiskovic/go-consumer-pattern/mocks/internal_/wine"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	//given
	barMock := mocks.NewBar(t)
	expected := "I want wine"
	barMock.EXPECT().IWantWine().Return(expected)
	//when
	sut := wine.New(barMock)
	//then
	assert.Equal(t, expected, sut.Order())
}
