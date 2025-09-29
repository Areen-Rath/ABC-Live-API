package handler;

import (
	"net/http"
	"ABC-Live-API/fetchers"
	"github.com/labstack/echo/v4"
);

type Data struct {
	Data	[]fetchers.News	`json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New();

	e.GET("/economic_times", func (c echo.Context) error {
		data := Data{fetchers.ETFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.GET("/moneycontrol", func (c echo.Context) error {
		data := Data{fetchers.MCFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.GET("/business_line", func (c echo.Context) error {
		data := Data{fetchers.BLFetcher()};
		return c.JSON(http.StatusOK, data);
	});

	e.ServeHTTP(w, r);
}