package main

import (
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func rubah(c echo.Context) error {
	t := c.FormValue("tarif")
	if (t == "") {
		return c.String(
			http.StatusBadRequest,
			"Tarif kosong")
	}

	_, err := strconv.ParseInt(t, 10, 16)
	if (err != nil) {
		return c.String(
			http.StatusBadRequest,
			"Tarif tidak layak")
	}

	_, erro := p.Exec(
		context.Background(),
		"INSERT INTO tarif (tarif) VALUES ($1)",
		t)
	if (erro != nil) {
		return c.String(
			http.StatusInternalServerError,
			"Tidak bisa memasukkan data")
	} else {
		return c.Redirect(
			http.StatusSeeOther,
			"/perubah")
	}
}
