package mocks

import (
	"github.com/BryanBD16/smallListApp/list"
)

// MockRepository is a mock implementation of the IRepository interface.
// It is used for testing purposes and does not interact with a real database.
type MockRepository struct {
	Elements []list.Element // Mock data storage for elements
	AddErr   error          // Simulated error for AddElement
	GetErr   error          // Simulated error for GetElements
	ClearErr error          // Simulated error for ClearList
}

// AddElement simulates adding an element to the repository.
func (m *MockRepository) AddElement(elem list.Element) error {
	// If AddErr is set, simulate an error
	if m.AddErr != nil {
		return m.AddErr
	}

	// Append the element to the mock storage
	m.Elements = append(m.Elements, elem)
	return nil
}

// GetElements simulates retrieving elements from the repository.
func (m *MockRepository) GetElements() ([]list.Element, error) {
	// If GetErr is set, simulate an error
	if m.GetErr != nil {
		return nil, m.GetErr
	}

	// Return the stored elements
	return m.Elements, nil
}

// ClearList simulates clearing the repository.
func (m *MockRepository) ClearList() error {
	// If ClearErr is set, simulate an error
	if m.ClearErr != nil {
		return m.ClearErr
	}

	// Clear the stored elements
	m.Elements = []list.Element{}
	return nil
}
