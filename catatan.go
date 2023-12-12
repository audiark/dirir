package main

import (
	"time"
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Model setiap deret pada tabel
type CatatanParkir struct {
	Picc	int
	Nama	string
	Waktu	time.Time
	Tipe	string
}

// Model untuk setiap halaman
type ModelCatatan struct {
	Catatan 	[]CatatanParkir
	Halaman 	uint64
	BisaSebelumnya	bool
	BisaSelanjutnya	bool
	Sebelumnya	uint64
	Selanjutnya	uint64
}

func catatan(c echo.Context) error {
	// Model-model dan penghitung model yang telah terisi 
	// Deretnya maksimal 12 buah
	var ca [12]CatatanParkir
	var i int

	// Query parameter "hal"
	// Jika kosong, maka akan diisikan 1
	h := c.QueryParam("hal")
	if h == "" {
		h = "1"
	}

	// Offset untuk query basis data
	// Dalam 1 halaman terdapat maksimal 12 deret data
	o, e := strconv.ParseUint(h, 10, 16)
	if e != nil {
		return c.String(
			http.StatusBadRequest,
			"hal tidak layak")
	} else {
		o = (o * 11) - 11
	}

	// Meng-query data
	r, er := p.Query(
		context.Background(),
		`SELECT picc, nama, waktu, tipe FROM catatan
		ORDER BY waktu DESC
		LIMIT 12 OFFSET $1`,
		o)
	if er != nil {
		return c.String(
			http.StatusInternalServerError,
			"Tidak dapat mendapatkan data")
	}

	// Mengisi model-model dengan data-data yang telah di-query
	// dan menghitung model yang telah terisi
	for i = 0; r.Next(); i++ {
		if (r.Scan(
		&(ca[i].Picc),
		&(ca[i].Nama),
		&(ca[i].Waktu),
		&(ca[i].Tipe)) != nil) {
			return c.String(
				http.StatusInternalServerError,
				"Tidak dapat me-scan data")
		}
	}

	// Mengekstrak halaman dari offset
	ha := (o + 11) / 11

	return c.Render(
		http.StatusOK,
		"catatan.html",
		ModelCatatan{
			ca[:i],
			ha,
			ha != 1 && i != 0,
			i == 12,
			ha - 1,
			ha + 1})
}
