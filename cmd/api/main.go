package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nathan/api-buku/internal/handler"
	"github.com/nathan/api-buku/storage"
)

func main() {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5432/postgres"
	conn, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v\n", err)
	}
	defer conn.Close()
	log.Println("Berhasil terhubung ke database")

	dbStorage := &storage.DBStorage{Conn: conn}
	bukuHandler := &handler.BukuHandler{Storage: dbStorage}
	http.HandleFunc("/api/buku", bukuHandler.GetAllBuku)
	port := ":8080"
	fmt.Printf("Server berjalan di port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
