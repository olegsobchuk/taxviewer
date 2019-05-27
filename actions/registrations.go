package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/olegsobchuk/taxviewer/models"
)

// RegistrationsNew default implementation.
func RegistrationsNew(c buffalo.Context) error {
	return c.Render(200, r.HTML("registrations/new.html"))
}

// RegistrationsCreate default implementation.
func RegistrationsCreate(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	tx, _ := c.Value("tx").(*pop.Connection)
	errval, some := user.Validate(tx)

	fmt.Printf("Err: %+v\n", errval)
	fmt.Printf("Err: %+v\n", some)
	pass := c.Param("password")
	passConfirm := c.Param("password_confirmation")
	if len(pass) < 6 || pass != passConfirm {
		c.Flash().Add("error", "Password incorrect")
		return c.Render(400, r.HTML("registrations/new.html"))
	}
	return c.Redirect(302, "newSessionsPath()")
}
