package router

import (
	"cic/site/pkg/models/gifts"
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func SpiritualGiftsRoutes(e *echo.Echo) {
	router := e.Group("/spiritual-gifts")
	router.GET("", spiritualGifts)
	router.POST("/submit", submitAnswers)
}

func spiritualGifts(c echo.Context) error {
	questions, err := gifts.GetQuestions()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "spiritual-gifts.html", questions)
}

func submitAnswers(c echo.Context) error {
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

	results, err := gifts.ProcessSpiritualGiftsResults(1)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf(
		"Top Gift: %v with %v points | Second Gift: %v with %v points | Third Gift: %v with %v points ",
		results.First.Name,
		results.FirstPoints,
		results.Second.Name,
		results.SecondPoints,
		results.Third.Name,
		results.ThirdPoints,
	))
}
