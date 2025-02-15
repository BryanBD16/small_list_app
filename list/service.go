package list

import (
	"encoding/json"
	"net/http"
)

type Service struct {
	l *list
	r IRepository // In Go, interfaces are inherently reference types, so you should not use pointers to interfaces (*IRepository)—only the interface type (IRepository) is needed.
}

func NewService(r IRepository) Service {
	s := Service{
		l: &list{},
		r: r,
	}

	return s
}

func (s *Service) Get(w http.ResponseWriter, req *http.Request) {
	elems, err := s.r.GetElements()

	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404 error
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(elems)

	if err != nil {
		http.Error(w, "Error occuring during elements encoding", http.StatusInternalServerError) //500 error
		return
	}
}

func (s *Service) Add(w http.ResponseWriter, req *http.Request) {
	var newElem Element
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newElem)
	if err != nil {
		http.Error(w, "Error when reading the data", http.StatusBadRequest) //400 error
		return
	}

	s.r.AddElement(newElem)

	w.WriteHeader(http.StatusCreated) //201 created
	w.Write([]byte("Element add successfully"))
}

func (s *Service) Clear(w http.ResponseWriter, req *http.Request) {
	// Call the repository method to clear the table
	err := s.r.ClearList()
	if err != nil {
		http.Error(w, "Failed to clear elements table", http.StatusInternalServerError) // 500 error
		return
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Elements table cleared successfully"))
}
