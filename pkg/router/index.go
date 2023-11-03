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
	var answer gifts.Answer
	err := c.Bind(&answer)
	if err != nil {
		return err
	}

	err = gifts.SubmitAnswer(answer.Answer, 1, answer.QuestionId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("Question %v: %v", answer.QuestionId, answer.Answer))
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

func DeptNews(c echo.Context) error {
	return c.Render(http.StatusOK, "dept-news.html", nil)
}

func KenyaTrip(c echo.Context) error {
	return c.Render(http.StatusOK, "kenya2024.html", nil)
}

func Dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.html", nil)
}

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "<h3><i>Greetings</i></h3>")
}
