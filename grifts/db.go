package grifts

import (
	"github.com/gobuffalo/vuerecipe/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		if err := models.DB.TruncateAll(); err != nil {
			return errors.WithStack(err)
		}

		category := &models.Category{
			Title: "The Beatles",
			Description:  "4 fun loving lads from Liverpool.",
		}
		if err := models.DB.Create(category); err != nil {
			return errors.WithStack(err)
		}
		posts := models.Posts{
			{Title: "John Lennon", Description: "Guitar"},
			{Title: "Paul McCartney", Description: "Bass"},
			{Title: "George Harrison", Description: "Guitar"},
			{Title: "Ringo Starr", Description: "Drums"},
		}
		for _, m := range posts {
			m.CategoryID = category.ID
			if err := models.DB.Create(&m); err != nil {
				return errors.WithStack(err)
			}
		}

		category = &models.Category{
			Title: "The Monkees",
			Description:  "4 fun loving lads assembled by a TV studio",
		}
		if err := models.DB.Create(category); err != nil {
			return errors.WithStack(err)
		}
		posts = models.Posts{
			{Title: "Mike Nesmith", Description: "Guitar"},
			{Title: "Davy Jones", Description: "Voice"},
			{Title: "Peter Tork", Description: "Guitar"},
			{Title: "Mikey Dolenz", Description: "Drums"},
		}
		for _, m := range posts {
			m.CategoryID = category.ID
			if err := models.DB.Create(&m); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})

})
