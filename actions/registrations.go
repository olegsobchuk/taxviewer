package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/olegsobchuk/taxviewer/models"
	"github.com/pkg/errors"
)

// RegistrationsNew default implementation.
func RegistrationsNew(c buffalo.Context) error {
	c.Set("user", &models.User{})
	return c.Render(200, r.HTML("registrations/new.html"))
}

// RegistrationsCreate default implementation.
func RegistrationsCreate(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}
	tx, _ := c.Value("tx").(*pop.Connection)
	validErrs, err := user.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if validErrs.HasAny() {
		c.Set("user", user)
		c.Set("errors", validErrs)
		c.Flash().Add("danger", "Please, check errors on the form!")
		return c.Render(200, r.HTML("registrations/new.html"))
	}
	c.Flash().Add("success", "Success! You can sign in now!")

	return c.Redirect(302, "newSessionsPath()")
}
