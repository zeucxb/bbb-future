package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestActions(t *testing.T) {
	g := Goblin(t)

	g.Describe("Make some action and return nil", func() {
		var err error

		g.It("Should make the run action", func() {
			err = runAction(nil)

			g.Assert(err).Equal(nil)
		})

		g.It("Should make the crowler action", func() {
			err = crowlerAction(nil)

			g.Assert(err).Equal(nil)
		})
	})
}
