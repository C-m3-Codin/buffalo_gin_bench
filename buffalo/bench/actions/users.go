// actions/users.go
package actions

import (
	"bench/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

// UsersResource is the resource for the User model
type UsersResource struct {
	buffalo.Resource
}

// List gets all Users. This function is mapped to the path
// GET /users
func (v UsersResource) List(c buffalo.Context) error {
	tx, _ := c.Value("tx").(*pop.Connection)
	users := []models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	return c.Render(200, r.JSON(users))
}

// Show gets the data for one User. This function is mapped to
// the path GET /users/{user_id}
func (v UsersResource) Show(c buffalo.Context) error {
	tx, _ := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := tx.Find(&user, c.Param("user_id")); err != nil {
		return err
	}
	return c.Render(200, r.JSON(user))
}

// Create adds a User to the DB. This function is mapped to the
// path POST /users
func (v UsersResource) Create(c buffalo.Context) error {
	tx, _ := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := tx.Create(&user); err != nil {
		return err
	}
	return c.Render(201, r.JSON(user))
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (v UsersResource) Update(c buffalo.Context) error {
	tx, _ := c.Value("tx").(*pop.Connection)
	user := models.User{}
	user.ID = uuid.FromStringOrNil(c.Param("user_id"))
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := tx.Update(&user); err != nil {
		return err
	}
	return c.Render(200, r.JSON(user))
}

// Destroy deletes a User from the DB. This function is mapped to
// the path DELETE /users/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	tx, _ := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := tx.Find(&user, c.Param("user_id")); err != nil {
		return err
	}
	if err := tx.Destroy(&user); err != nil {
		return err
	}
	return c.Render(200, nil)
}
