package router

import (
	"cic/site/pkg/models/gifts"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BaseRoutes(e *echo.Echo) {
	e.GET("/", indexRoute)
	e.GET("/htmx-test", test)
	e.GET("/test/q", questionsTest)
	e.GET("/test/g", giftsTest)
	e.GET("/dashboard", dashboard)
	e.GET("/kenya2024", kenyaTrip)
	e.GET("/dept-news", deptNews)
}

func indexRoute(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func questionsTest(c echo.Context) error {
	questions, err := gifts.GetQuestions()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "questions.html", questions)
}

func giftsTest(c echo.Context) error {
	g, err := gifts.GetGifts()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "gifts.html", g)
}

func deptNews(c echo.Context) error {
	return c.Render(http.StatusOK, "dept-news.html", nil)
}

func kenyaTrip(c echo.Context) error {
	return c.Render(http.StatusOK, "kenya2024.html", nil)
}

func dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.html", nil)
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "<h3><i>Greetings</i></h3>")
}
