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
			Title: "Laptop",
			Description:  "this is category laptop",
		}
		if err := models.DB.Create(category); err != nil {
			return errors.WithStack(err)
		}
		posts := models.Posts{
			{Title: "Laptop new 1", Description: "Guitar"},
			{Title: "Laptop new 2", Description: "Bass"},
			{Title: "Laptop new 3", Description: "Guitar"},
			{Title: "Laptop new 4", Description: "Drums"},
		}
		for _, m := range posts {
			m.CategoryID = category.ID
			if err := models.DB.Create(&m); err != nil {
				return errors.WithStack(err)
			}
		}

		category = &models.Category{
			Title: "Mobile",
			Description:  "this is mobile",
		}
		if err := models.DB.Create(category); err != nil {
			return errors.WithStack(err)
		}
		posts = models.Posts{
			{Title: "Mobile Mike Nesmith", Description: "Guitar"},
			{Title: "Mobile Davy Jones", Description: "Voice"},
			{Title: "Mobile Peter Tork", Description: "Guitar"},
			{Title: "Mobile Mikey Dolenz", Description: "Drums"},
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
