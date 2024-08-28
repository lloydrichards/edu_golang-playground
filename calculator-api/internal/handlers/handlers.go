package handlers

import (
	"encoding/json"
	"net/http"
)

type Body struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := body.Value1 + body.Value2

	writeResponse(w, float32(result))
}

func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := body.Value1 - body.Value2

	writeResponse(w, float32(result))
}

func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := body.Value1 * body.Value2

	writeResponse(w, float32(result))
}

func HandleDivide(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if body.Value2 == 0 {
		http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		return
	}

	result := body.Value1 / body.Value2

	writeResponse(w, float32(result))
}

// --- Helper functions ---

func getBody(r *http.Request) (Body, error) {
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}

func writeResponse(
	w http.ResponseWriter,
	result float32,
) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float32{"result": result})

}
