package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/olegsobchuk/taxviewer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
