package rest

import (
	"encoding/json"
	"net/http"

	api "github.com/Kronverskiy/habit/api/gen"
)

type HabitHandler struct {
}

func NewHabitHandler() *HabitHandler {
	return &HabitHandler{}
}

var _ api.ServerInterface = (*HabitHandler)(nil)

func (h *HabitHandler) GetHabitByID(w http.ResponseWriter, r *http.Request, habitID int) {
	testResult := api.Habit{
		Id:   habitID,
		Name: "Test habit pipipupu",
	}

	jsonRes, err := json.MarshalIndent(testResult, "", "  ")
	if err != nil {
		http.Error(w, "some error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
