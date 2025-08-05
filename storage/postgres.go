package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nathan/api-buku/internal/model"
)

type DBStorage struct {
	Conn *pgxpool.Pool
}

func (s *DBStorage) GetAllBuku(ctx context.Context) ([]model.Buku, error) {
	query := "SELECT id, judul, penulis, tahun FROM buku ORDER BY id"
	rows, err := s.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var daftarBuku []model.Buku
	for rows.Next() {
		var b model.Buku
		if err := rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
		}
		daftarBuku = append(daftarBuku, b)
	}
	return daftarBuku, nil
}
