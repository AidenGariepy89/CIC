package main

import (
	"cic/site/pkg/db"
	"cic/site/pkg/models/gifts"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const template_files string = "views"

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	templates, err := template.ParseGlob(path.Join(template_files, "*.html"))
	if err != nil {
		log.Fatalf("Error loading templates: %v\n", err)
	}

	err = db.InitDb("file:tmp/quiz.db")
	if err != nil {
		log.Fatalf("Error creating db: %v\n", err)
	}

	err = test()
	if err != nil {
		log.Fatalf("Error testing: %v\n", err)
	}

	e := echo.New()

	e.Renderer = &TemplateRenderer{
		template: templates,
	}

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":42069"))
}

func test() error {
	questions, err := gifts.GetQuestions()
	if err != nil {
		return err
	}

	for i := 0; i < len(*questions); i++ {
		question := &(*questions)[i]

		log.Printf("Question: %v: %c | %v\n", question.Id, question.Gift, question.Content)
	}

	return nil
}
