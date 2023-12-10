package main

import (
	"context"
	"fmt"
	"os"
	"io"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

var p *pgxpool.Pool

func main() {
	var er error

  	// Echo, alamat, dan hubungan database
  	e 	:= echo.New()
  	a 	:= ":" + os.Getenv("PORT")
	p, er	= pgxpool.New(
			context.Background(),
			os.Getenv("DATABASE_URL"))
	
	if er != nil {
		fmt.Fprintf(
			os.Stderr,
			"Tidak bisa terhubung dengan basis data: %v\n",
			er)

		os.Exit(1)
	}
	
	defer p.Close()
	
	e.Renderer = &MyTemplate {
		templates:
			template.Must(template.ParseGlob("tmpl/*.html")),
	}

	// Middleware
	e.Use(middleware.Logger())
  	e.Use(middleware.Recover())

  	// Rute-rutenya
  	e.GET("/", beranda)
	e.GET("/catatan", catatan)
	e.GET("/perubah", perubah)
	e.GET("/tarif", tarif)
	e.GET("/jumlah", jumlah)
	e.POST("/masuk", masuk)
	e.POST("/keluar", keluar)
	e.POST("/rubah", rubah)

  	// Start server
	e.Logger.Fatal(e.Start(a))
}

type MyTemplate struct {
	templates *template.Template
}

func (t *MyTemplate) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

