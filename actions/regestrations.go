package actions

import "github.com/gobuffalo/buffalo"

// RegestrationsNew default implementation.
func RegestrationsNew(c buffalo.Context) error {
	return c.Render(200, r.HTML("regestrations/new.html"))
}

// RegestrationsCreate default implementation.
func RegestrationsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("regestrations/create.html"))
}

