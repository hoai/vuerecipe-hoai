package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/vuerecipe/models"
	"github.com/pkg/errors"
	"fmt"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Post)
// DB Table: Plural (posts)
// Resource: Plural (Posts)
// Path: Plural (/posts)
// View Template Folder: Plural (/templates/posts/)

// PostsResource is the resource for the Post model
type PostsResource struct{
  buffalo.Resource
}

// List gets all Posts. This function is mapped to the path
// GET /posts
func (v PostsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx := c.Value("tx").(*pop.Connection)

  posts := &models.Posts{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

	// Retrieve all Members from the DB
	if err := q.Where("category_id = ?", c.Param("category_id")).All(posts); err != nil {
		return errors.WithStack(err)
	}

  // Add the paginator to the headers so clients know how to paginate.
	c.Response().Header().Set("X-Pagination", q.Paginator.String())

	return c.Render(200, r.JSON(posts))
}

// Show gets the data for one Post. This function is mapped to
// the path GET /posts/{post_id}
func (v PostsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx := c.Value("tx").(*pop.Connection)

  // Allocate an empty Post
  post := &models.Post{}

  // To find the Post the parameter post_id is used.
  if err := tx.Find(post, c.Param("post_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.JSON(post))
}

// New renders the form for creating a new Post.
// This function is mapped to the path GET /posts/new
func (v PostsResource) New(c buffalo.Context) error {
  return c.Error(404, errors.New("not available"))
}
// Create adds a Post to the DB. This function is mapped to the
// path POST /posts
func (v PostsResource) Create(c buffalo.Context) error {
  // Allocate an empty Post
  post := &models.Post{}

  // Bind post to the html form elements
  if err := c.Bind(post); err != nil {
    return errors.WithStack(err)
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(post)
  if err != nil {
    return errors.WithStack(err)
  }

  if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(201, r.JSON(post))
}

// Edit renders a edit form for a Post. This function is
// mapped to the path GET /posts/{post_id}/edit
func (v PostsResource) Edit(c buffalo.Context) error {
	return c.Error(404, errors.New("not available"))
}
// Update changes a Post in the DB. This function is mapped to
// the path PUT /posts/{post_id}
func (v PostsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx := c.Value("tx").(*pop.Connection)
  
  // Allocate an empty Post
  post := &models.Post{}

  if err := tx.Where("category_id = ?", c.Param("category_id")).Find(post, c.Param("post_id")); err != nil {
		return errors.WithStack(err)
	}
  // Bind Post to the html form elements
  if err := c.Bind(post); err != nil {
    return errors.WithStack(err)
  }

  verrs, err := tx.ValidateAndUpdate(post)
  if err != nil {
    return errors.WithStack(err)
  }

  if verrs.HasAny() {
		// Render errors as JSON
		return c.Render(400, r.JSON(verrs))
	}

	return c.Render(200, r.JSON(post))
}

// Destroy deletes a Post from the DB. This function is mapped
// to the path DELETE /posts/{post_id}
func (v PostsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx := c.Value("tx").(*pop.Connection)

  // Allocate an empty Post
  post := &models.Post{}

  // To find the Post the parameter post_id is used.
  if err := tx.Where("category_id = ?", "category_id").Find(post, c.Param("post_id")); err != nil {
		return c.Error(404, err)
	}

  if err := tx.Destroy(post); err != nil {
    return errors.WithStack(err)
  }

  return c.Render(200, r.JSON(post))
}
