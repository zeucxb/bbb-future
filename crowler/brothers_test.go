package crowler

import (
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestFormatBirthdate(t *testing.T) {
	g := Goblin(t)

	g.Describe("Format Birthdate", func() {
		var formatedBirthdate time.Time
		var err error

		g.It("Should format a string in the format DD/MM/YYYY to time.Time", func() {
			assertationDate := time.Date(1997, time.May, 16, 0, 0, 0, 0, time.UTC)

			formatedBirthdate, err = formatBirthdate("16/05/1997")
			if err != nil {
				t.Error(err)
			}
			g.Assert(formatedBirthdate).Equal(assertationDate)
		})

		g.It("Should format a string in the format DD/MM/YYYY - DD/MM/YYYY to time.Time", func() {
			assertationDate := time.Date(1997, time.May, 16, 0, 0, 0, 0, time.UTC)

			formatedBirthdate, err = formatBirthdate("16/05/1997 - 31/12/2100")
			if err != nil {
				t.Error(err)
			}
			g.Assert(formatedBirthdate).Equal(assertationDate)
		})

		g.It("Should format a string in the format (YYYY) to time.Time", func() {
			assertationDate := time.Date(1997, time.January, 1, 0, 0, 0, 0, time.UTC)

			formatedBirthdate, err = formatBirthdate("16 (1997)")

			g.Assert(err).Equal(nil)
			g.Assert(formatedBirthdate).Equal(assertationDate)
		})
	})
}

func TestSaveBrothersData(t *testing.T) {
	g := Goblin(t)

	g.Describe("Save Brothers Data", func() {
		var err error

		g.It("Should save the data and return without error", func() {
			err = SaveBrothersData()

			g.Assert(err).Equal(nil)
		})
	})
}
