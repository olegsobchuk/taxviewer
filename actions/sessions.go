package actions

import (
	"database/sql"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/olegsobchuk/taxviewer/models"
	"github.com/pkg/errors"
)

// SessionsNewHandler is a SignIn page.
func SessionsNewHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("sessions/new.html"))
}

// SessionsCreateHandler is a SignIn page.
func SessionsCreateHandler(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Where("email = ?", strings.ToLower(user.Email)).First(user)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			c.Flash().Add("danger", "Email or password incorrect.")
			return c.Redirect(302, "newSessionsPath()") // TODO: should be changed
		}
		return errors.WithStack(err)
	}

	isValid := user.CheckPassword()
	if isValid {
		// TODO: set session
		// redirect in system
		c.Flash().Add("success", "Success session.")
	} else {
		c.Flash().Add("danger", "Email or password incorrect.")
	}
	return c.Redirect(302, "newSessionsPath()")
}
