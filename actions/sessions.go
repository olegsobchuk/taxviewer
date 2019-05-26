package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/olegsobchuk/taxviewer/models"
)

// SessionsNewHandler is a SignIn page.
func SessionsNewHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("sessions/new.html"))
}

// SessionsCreateHandler is a SignIn page.
func SessionsCreateHandler(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	user.CheckPassword(c.Param("password"))
	return c.Redirect(302, "newSessionsPath()")
}
