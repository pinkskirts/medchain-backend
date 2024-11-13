package controler

import (
	"encoding/json"
	"net/http"

	"github.com/pinkskirts/medchain-backend/usecase"
)

type PrescriptionController struct {
	usecase *usecase.PrescriptionUseCase
}

func NewPrescriptionController(usecase *usecase.PrescriptionUseCase) *PrescriptionController {
	return &PrescriptionController{
		usecase: usecase,
	}
}

func (p *PrescriptionController) GetPrescription(w http.ResponseWriter, r *http.Request) {
	dtos, err := p.usecase.FetchPrescription()
	if err != nil {
		http.Error(w, "Failed to get prescriptions", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dtos)
	if err != nil {
		http.Error(w, "Failed to get prescriptions", http.StatusInternalServerError)
		return
	}
}

func (p *PrescriptionController) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	var command *usecase.CreatePrescriptionRequest
	err := json.NewDecoder(r.Body).Decode(command)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := p.usecase.CreatePrescription(command)
	if err != nil {
		http.Error(w, "failed to create prescription", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message":             "Prescription created successfully",
		"prescriptionAddress": id,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
