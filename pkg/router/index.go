package router

import (
	"cic/site/pkg/models/gifts"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func SpiritualGifts(c echo.Context) error {
	questions, err := gifts.GetQuestions()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "spiritual-gifts.html", questions)
}

func SubmitAnswers(c echo.Context) error {
	// var answer gifts.Answer
	// err := c.Bind(&answer)
	// if err != nil {
	//     return err
	// }

	var val string
	val = c.Param("content")

	return c.String(http.StatusOK, fmt.Sprintf("Go it: '%#v'", val))
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
