package list

import "fmt"

type listCreator interface {
	AddElement(elem Element) //no need to import since the Element type struct is declared inside the same package (list)
}

type listGetter interface {
	GetElements()
}

type list struct {
	elements []Element
}

func (l *list) AddElement(elem Element) error {
	l.elements = append(l.elements, elem)

	return nil
}

func (l *list) GetElements() ([]Element, error) {
	if l.elements == nil {
		return nil, fmt.Errorf("the list has no elements")
	}
	return l.elements, nil
}
