package router

import (
	"cic/site/pkg/models/gifts"
	"cmp"
	"net/http"
	"slices"
	"strconv"
	"strings"

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

func SubmitAnswers(c echo.Context) error {
	params, err := c.FormParams()
	if err != nil {
		return err
	}

	answers := []gifts.Answer{}

	for param := range params {
		id, err := strconv.Atoi(strings.Split(param, "-")[1])
		if err != nil {
			return err
		}

		answer, err := strconv.Atoi(c.FormValue(param))
		if err != nil {
			return err
		}

		answers = append(answers, gifts.Answer{
			UserId:     1,
			QuestionId: id,
			Answer:     answer,
		})
	}

	slices.SortFunc(answers, func(a, b gifts.Answer) int {
		return cmp.Compare(a.QuestionId, b.QuestionId)
	})

	for _, answer := range answers {
		err := gifts.SubmitAnswer(answer.Answer, answer.UserId, answer.QuestionId)
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "Working?")
}
