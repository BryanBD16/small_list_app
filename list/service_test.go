package list

//run the command go test -v to see the detail of all the test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRepository is a mock implementation of the IRepository interface.
// It is used for testing purposes and does not interact with a real database.
type MockRepository struct {
	Elements []Element // Mock data storage for elements
	AddErr   error     // Simulated error for AddElement
	GetErr   error     // Simulated error for GetElements
	ClearErr error     // Simulated error for ClearList
}

// AddElement simulates adding an element to the repository.
func (m *MockRepository) AddElement(elem Element) error {
	// If AddErr is set, simulate an error
	if m.AddErr != nil {
		return m.AddErr
	}

	// Append the element to the mock storage
	m.Elements = append(m.Elements, elem)
	return nil
}

// GetElements simulates retrieving elements from the repository.
func (m *MockRepository) GetElements() ([]Element, error) {
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
	m.Elements = []Element{}
	return nil
}

// TestService_Get tests the Get method of the Service
func TestService_Get(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		Elements: []Element{ // Use the list.Element struct
			{Name: "Element 1", Description: "First element"},
			{Name: "Element 2", Description: "Second element"},
		},
	}

	service := NewService(mockRepo) // Use the list.NewService function
	req, err := http.NewRequest("GET", "/elements", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest.NewRecorder to capture the response
	rr := httptest.NewRecorder()

	// Act
	service.Get(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Element 1")
	assert.Contains(t, rr.Body.String(), "Element 2")
}

// TestService_Add tests the Add method of the Service
func TestService_Add(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{}
	service := NewService(mockRepo) // Use the list.NewService function

	// Create a new element for the request body
	newElem := Element{ // Use the list.Element struct
		Name:        "New Element",
		Description: "This is a new element",
	}

	// Convert the new element to JSON
	reqBody, err := json.Marshal(newElem)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/elements", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest.NewRecorder to capture the response
	rr := httptest.NewRecorder()

	// Act
	service.Add(rr, req)

	// Assert
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Element add successfully", rr.Body.String())
	assert.Len(t, mockRepo.Elements, 1) // Verify the element is added
	assert.Equal(t, "New Element", mockRepo.Elements[0].Name)
}

// TestService_Clear tests the Clear method of the Service
func TestService_Clear(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{
		Elements: []Element{ // Use the list.Element struct
			{Name: "Element 1", Description: "First element"},
		},
	}
	service := NewService(mockRepo) // Use the list.NewService function

	req, err := http.NewRequest("POST", "/clear", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest.NewRecorder to capture the response
	rr := httptest.NewRecorder()

	// Act
	service.Clear(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Elements table cleared successfully", rr.Body.String())
	assert.Len(t, mockRepo.Elements, 0) // Elements should be cleared
}
