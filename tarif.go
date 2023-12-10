package main

import (
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func tarif(c echo.Context) error {
	var t int
	err := p.QueryRow(
		context.Background(),
		"SELECT tarif_terkini()").Scan(&t)
	if (err != nil) {
		return c.String(
			http.StatusInternalServerError,
			"Tidak bisa meng-query atau me-scan")
	} else {
		return c.String(
			http.StatusOK,
			strconv.FormatInt(int64(t), 10))
	}
}
