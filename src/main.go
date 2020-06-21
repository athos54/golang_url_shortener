package main

import (
	"net/http"

	m "./models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Init shortener")

	m.Connect()

	e := echo.New()

	e.GET("/:short_path", func(c echo.Context) error {
		us := m.UrlShorten{}
		us.ShortPath = c.Request().RequestURI

		log.Info(c.Request().RequestURI)

		m.SearchURL(&us)

		if us.LongURL == "" {
			return c.JSON(400, "")
		}

		return c.Redirect(http.StatusSeeOther, us.LongURL)
	})

	e.POST("/", func(c echo.Context) error {
		us := m.UrlShorten{}

		if err := c.Bind(&us); err != nil {
			log.Error("error bind")
			return c.JSON(400, "")
		}

		log.Info(us)

		if err := m.AppendURL(&us); err != nil {
			log.Error("Error append: " + err.Error())
			return c.JSON(400, "")
		}

		return c.JSON(200, "")

	})

	e.Logger.Fatal(e.Start(":1323"))
}
