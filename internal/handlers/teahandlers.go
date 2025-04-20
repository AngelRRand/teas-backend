package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"strconv"
	"sync"
	"teas/internal/models"
	"teas/internal/utils"
)

var (
	teaData     []models.TeaCategory
	teaDataOnce sync.Once
	teaDataErr  error
)

func LoadTeaData(path string) error {
	teaDataOnce.Do(func() {
		data, err := os.ReadFile(path)
		if err != nil {
			teaDataErr = err
			return
		}

		var categories []models.TeaCategory
		if err := json.Unmarshal(data, &categories); err != nil {
			teaDataErr = err
			return
		}

		teaData = categories
	})
	return teaDataErr
}

func TeaAllData(w http.ResponseWriter, r *http.Request) {
	if teaDataErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, "No se pudieron cargar los datos")
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, teaData)
}

func ListCategories(w http.ResponseWriter, r *http.Request) {
	if teaDataErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, "No se pudieron cargar los datos")
		return
	}
	categories := make([]string, 0, len(teaData))
	for _, te := range teaData {
		categories = append(categories, te.Category)
	}
	utils.WriteJSONResponse(w, http.StatusOK, categories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	if teaDataErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, "No se pudieron cargar los datos")
		return
	}

	idParam := chi.URLParam(r, "id")
	num, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "ID de categoría inválido")
		return
	}

	if num < 0 || num >= len(teaData) {
		utils.WriteError(w, http.StatusNotFound, "Categoría no encontrada")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, teaData[num])
}
