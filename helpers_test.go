package queryBuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgsCountNegative(t *testing.T) {
	actual := ArgsCount(-10)
	expected := ""
	assert.Equal(t, expected, actual, "")
}

func TestArgsCount0(t *testing.T) {
	inValues := []int{}
	actual := ArgsCount(len(inValues))
	expected := ""
	assert.Equal(t, expected, actual, "")
}

func TestArgsCount1(t *testing.T) {
	inValues := []string{"1"}
	actual := ArgsCount(len(inValues))
	expected := "?"
	assert.Equal(t, expected, actual, "")
}

func TestArgsCount5(t *testing.T) {
	inValues := []int{1}
	actual := ArgsCount(len(inValues))
	expected := "?"
	assert.Equal(t, expected, actual, "")
}
