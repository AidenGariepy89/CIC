package main

import (
	"cic/site/pkg/db"
	"cic/site/pkg/router"
	"html/template"
	"io"
	"log"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

    "github.com/gotailwindcss/tailwind/twembed"
    "github.com/gotailwindcss/tailwind/twhandler"

)

const template_files string = "views"

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	// Parse templates
	templates, err := template.ParseGlob(path.Join(template_files, "*.html"))
	if err != nil {
		log.Fatalf("Error loading templates: %v\n", err)
	}

	// Initialize Database
	err = db.InitDb("file:tmp/quiz.db")
	if err != nil {
		log.Fatalf("Error creating db: %v\n", err)
	}

	// Setup Echo server
	e := echo.New()

	e.Renderer = &TemplateRenderer{
		template: templates,
	}

	// Middlware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    // Static files
    e.Static("/js", "public/js")
    e.Static("/css", "public/css")

	// Routing
	e.GET("/", router.Index)
	e.GET("/q", router.Questions)

	// Serve
	e.Logger.Fatal(e.Start(":42069"))
}
