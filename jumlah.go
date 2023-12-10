package main

import (
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func jumlah(c echo.Context) error {
	var j int
	err := p.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM terparkir").Scan(&j)

	if (err != nil) {
		return c.String(
			http.StatusInternalServerError,
			"Tidak bisa meng-query atau me-scan")
	} else {
		return c.String(http.StatusOK, strconv.Itoa(j))
	}
}
