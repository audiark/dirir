package main

import (
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func masuk(c echo.Context) error {
	pi := c.FormValue("picc")
	n := c.FormValue("nama")
	if (pi == "" || n == "") {
		return c.String(
			http.StatusBadRequest,
			"Uid atau nama kosong")
	}

	_, err := strconv.ParseInt(pi, 10, 16)
	if (err != nil) {
		return c.String(
			http.StatusBadRequest,
			"Uid tidak layak")
	}

	_, erro := p.Exec(
		context.Background(),
		"INSERT INTO terparkir VALUES ($1, $2)",
		pi,
		n)
	if (erro != nil) {
		return c.String(
			http.StatusInternalServerError,
			"Tidak bisa memasukkan data")
	} else {
		return c.String(http.StatusOK, "Masuk")
	}
}
