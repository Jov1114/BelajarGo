package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nathan/api-buku/storage"
)

type BukuHandler struct {
	Storage *storage.DBStorage
}

func (h *BukuHandler) GetAllBuku(w http.ResponseWriter, r *http.Request) {
	buku, err := h.Storage.GetAllBuku(r.Context())
	if err != nil {
		http.Error(w, "Gagal mengambil data buku", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buku)
}
