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

	// get the request body
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add the values
	result := body.Value1 + body.Value2

	writeResponse(w, float32(result))
}

func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "subtract"}`))
}

func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "multiply"}`))
}

func HandleDivide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "divide"}`))
}

func HandleSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "sum"}`))
}

// --- Helper functions ---

func writeResponse(
	w http.ResponseWriter,
	result float32,
) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float32{"result": result})

}
