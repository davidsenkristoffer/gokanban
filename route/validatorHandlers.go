package route

import (
	"gokanban/components"
	"gokanban/helpers"

	"github.com/labstack/echo/v4"
)

func validateTitle(c echo.Context) error {
	title := c.QueryParam("title")
	errors := helpers.ValidateTitle(title)

	cmp := components.DisplayErrors(errors)

	return View(c, cmp)
}

func validateDescription(c echo.Context) error {
	description := c.QueryParam("description")
	errors := helpers.ValidateDescription(description)

	cmp := components.DisplayErrors(errors)
	return View(c, cmp)
}

func validateEstimatedtime(c echo.Context) error {
	estimatedtime := c.QueryParam("estimatedtime")
	errors := helpers.ValidateTime(estimatedtime)

	cmp := components.DisplayErrors(errors)
	return View(c, cmp)
}

func validateSpenttime(c echo.Context) error {
	estimatedtime := c.QueryParam("spenttime")
	errors := helpers.ValidateTime(estimatedtime)

	cmp := components.DisplayErrors(errors)
	return View(c, cmp)
}
