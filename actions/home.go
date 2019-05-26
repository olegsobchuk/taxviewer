package actions

import "github.com/gobuffalo/buffalo"

// RoutesHandler is a default handler to serve up a home page.
func RoutesHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("routes.html"))
}

// HomeHandler is a root and home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("names", []string{"Oleh", "Ehor"})
	return c.Render(200, r.HTML("home.html"))
}
