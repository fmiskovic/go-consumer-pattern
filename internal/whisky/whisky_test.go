package whisky_test

import (
	"testing"

	"github.com/fmiskovic/go-consumer-pattern/internal/whisky"
	mocks "github.com/fmiskovic/go-consumer-pattern/mocks/internal_/whisky"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	barMock := mocks.NewBar(t)
	expected := "I want whisky"
	barMock.EXPECT().IWantWhisky().Return(expected)
	//when
	sut := whisky.New(barMock)
	//then
	assert.Equal(t, expected, sut.Order())
}
