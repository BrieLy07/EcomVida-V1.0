package handlers

import (
	"encoding/json"
	"net/http"

	"product-image-service/utils"

	"github.com/gorilla/mux"
)

func SubirImagen(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("imagen")
	if err != nil {
		http.Error(w, "Imagen no enviada", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename
	contentType := header.Header.Get("Content-Type")

	err = utils.SubirAminio(file, filename, contentType)
	if err != nil {
		http.Error(w, "Error al subir imagen", http.StatusInternalServerError)
		return
	}

	utils.GuardarEnCache(filename)

	json.NewEncoder(w).Encode(map[string]string{"ok": "true", "archivo": filename})
}

func ObtenerImagen(w http.ResponseWriter, r *http.Request) {
	nombre := mux.Vars(r)["nombre"]

	// Verificar si está en Redis
	if url := utils.ObtenerDeCache(nombre); url != "" {
		json.NewEncoder(w).Encode(map[string]string{"url": url})
		return
	}

	// Si no está, generar URL firmada y guardar en Redis
	url, err := utils.GenerarURLImagen(nombre)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusNotFound)
		return
	}

	utils.GuardarEnCache(nombre)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
