package main

import (
	"strconv"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func keluar(c echo.Context) error {
	pi := c.FormValue("picc")
	if (pi == "") {
		return c.String(
			http.StatusBadRequest,
			"Uid kosong")
	}

	_, err := strconv.ParseInt(pi, 10, 16)
	if (err != nil) {
		return c.String(
			http.StatusBadRequest,
			"Uid tidak layak")
	}

	_, erro := p.Exec(
		context.Background(),
		"DELETE FROM terparkir WHERE picc = $1",
		pi)
	if (erro != nil) {
		return c.String(
			http.StatusInternalServerError,
			"Tidak bisa menghapus entri")
	} else {
		return c.String(http.StatusOK, "Keluar")
	}
}