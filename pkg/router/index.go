package router

import (
	"cic/site/pkg/models/gifts"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Questions(c echo.Context) error {
	questions, err := gifts.GetQuestions()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "questions.html", questions)
}

func Gifts(c echo.Context) error {
	g, err := gifts.GetGifts()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "gifts.html", g)
}

func SubmitAnswers(c echo.Context) error {
	err := gifts.SubmitAnswer(0, 1, 1)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "test answer added to database")
}
