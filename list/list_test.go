package list

// to run the test you have to be in the right directory (list not SmallListApp)

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// the go test command is used to run the test

func Test_AddElement(t *testing.T) {
	//Arrange
	l := &list{}
	elem := Element{
		Name:        "test",
		Description: "it is a test element",
	}

	//Act
	err := l.AddElement(elem)

	//Assert
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, 1, len(l.elements), "should have only one element")
}

func Test_GetElement(t *testing.T) {
	//Arrange
	l := &list{}
	elem := Element{
		Name:        "test",
		Description: "it is a test",
	}

	//Act
	l.AddElement(elem)

	list, err := l.GetElements()

	if err != nil {

	}
	if len(list) == 1 {

	}

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(list))
}

func Test_GetElementEmpty(t *testing.T) {
	l := &list{}
	_, err := l.GetElements()

	assert.Error(t, err)
}
