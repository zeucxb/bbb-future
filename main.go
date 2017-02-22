package main

import (
	"fmt"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/celrenheit/spider"
)

// BBBSpider scrape wikipedia's page for BBB
// It is defined below in the init function
var BBBSpider spider.Spider

// func main() {
// 	// Create a new scheduler
// 	scheduler := spider.NewScheduler()

// 	// Register the spider to be scheduled every 15 seconds
// 	scheduler.Add(schedule.Every(1*time.Millisecond), BBBSpider)

// 	// Start the scheduler
// 	scheduler.Start()

// 	// Exit 5 seconds later to let time for the request to be done.
// 	// Depends on your internet connection
// 	<-time.After(65 * time.Second)
// }

type originType struct {
	City    string `bson:"origin"`
	State   string `bson:"state"`
	Country string `bson:"country"`
}

type brotherType struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Birthdate  string        `bson:"birthdate"`
	Origin     originType    `bson:"origin"`
	Occupation string        `bson:"occupation"`
}

type brothersList []interface{}

func main() {
	var newBrothers brothersList

	BBBSpider = spider.Get("https://pt.wikipedia.org/wiki/Lista_de_participantes_do_Big_Brother_Brasil", func(ctx *spider.Context) error {
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

					newBrother := brotherType{
						ID:        bson.NewObjectId(),
						Name:      brother.Eq(0).Text(),
						Birthdate: brother.Eq(1).Text(),
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

	ctx, err := BBBSpider.Setup(nil)
	if err != nil {
		panic(err)
	}

	err = BBBSpider.Spin(ctx)
	if err != nil {
		panic(err)
	}
}
