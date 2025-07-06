package handlers

import (
	"address-service/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteDireccionHandler(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idStr := params["addressId"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = repo.Eliminar(id)
		if err != nil {
			http.Error(w, "Error al eliminar dirección", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
