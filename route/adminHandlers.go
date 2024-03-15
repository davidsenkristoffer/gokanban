package route

import (
	"gokanban/components"
	"gokanban/structs"

	"github.com/labstack/echo/v4"
)

func getAdminIndex(c echo.Context) error {

	db := c.(*kanbanContext).db
	ts := c.(*kanbanContext).ts

	tvm, err := ts.GetTags(db)
	if err != nil {
		tvm = make([]structs.TagViewModel, 0)
	}

	adm := &structs.AdminViewModel{
		Tags: tvm,
	}

	cmp := components.Admin(*adm)
	return View(c, cmp)
}

func getTags(c echo.Context) error {
	db := c.(*kanbanContext).db
	ts := c.(*kanbanContext).ts

	tvm, err := ts.GetTags(db)
	if err != nil {
		return c.NoContent(500)
	}

	cmp := components.AdminTaglist(tvm)
	return View(c, cmp)
}

func createTag(c echo.Context) error {
	db := c.(*kanbanContext).db
	ts := c.(*kanbanContext).ts

	tvm := &structs.TagViewModel{
		Label: c.FormValue("title"),
		Color: c.FormValue("tagcolor"),
	}
	t, err := tvm.ToModel()
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(400)
	}

	_, err = ts.CreateTag(db, t)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(500)
	}

	c.Response().Header().Set("HX-Trigger", "tags-updated")
	return c.NoContent(201)
}
