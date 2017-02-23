package crowler

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/celrenheit/spider"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type originType struct {
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
}

type brotherType struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Birthdate  time.Time     `bson:"birthdate"`
	Origin     originType    `bson:"origin"`
	Occupation string        `bson:"occupation"`
	Edition    int           `bson:"edition"`
}

type brothersList []interface{}

var newBrothers brothersList

var bbbSpider = spider.Get("https://pt.wikipedia.org/wiki/Lista_de_participantes_do_Big_Brother_Brasil", func(ctx *spider.Context) error {
	fmt.Println(time.Now())
	// Execute the request
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	// Get goquery's html parser
	htmlparser, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	summary := htmlparser.Find(".wikitable > tbody")

	for i := 0; i < summary.Length(); i++ {
		brothers := summary.Eq(i).Find(".wikitable > tbody > tr")

		for j := 0; j < brothers.Length(); j++ {
			brother := brothers.Eq(j).Find(".wikitable > tbody > tr > td")

			if brother.Length() > 0 {
				dateRE := regexp.MustCompile(`\(\d{4}\)`)
				birthdateString := dateRE.FindString(brother.Eq(1).Text())
				if birthdateString != "" {
					birthdateString = strings.Trim(birthdateString, "()")
					birthdateString = fmt.Sprintf("01/01/%s", birthdateString)
				} else {
					birthdateString = strings.Split(brother.Eq(1).Text(), "-")[0]
					birthdateString = strings.Trim(birthdateString, " ")
				}

				birthdate, err := time.Parse("02/01/2006", birthdateString)
				if err != nil {
					return err
				}

				newBrother := brotherType{
					ID:        bson.NewObjectId(),
					Name:      brother.Eq(0).Text(),
					Birthdate: birthdate,
					Edition:   i + 1,
				}

				var origin []string

				if i < 12 {
					newBrother.Occupation = brother.Eq(2).Text()

					origin = strings.Split(brother.Eq(3).Text(), ",")
				} else {
					newBrother.Occupation = brother.Eq(3).Text()

					origin = strings.Split(brother.Eq(2).Text(), ",")
				}

				if len(origin) > 1 {
					newBrother.Origin.City = strings.Trim(origin[0], " ")
					newBrother.Origin.State = strings.Trim(origin[1], " ")
					newBrother.Origin.Country = "Brasil"
				} else {
					newBrother.Origin.Country = origin[0]
				}

				newBrothers = append(newBrothers, newBrother)
			}
		}
	}

	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		return err
	}

	defer s.Close()

	err = s.DB("bbb").C("brothers").Insert(newBrothers...)

	return err
})

func SaveBrothersData() (err error) {
	ctx, err := bbbSpider.Setup(nil)
	if err != nil {
		return
	}

	err = bbbSpider.Spin(ctx)

	return

}
