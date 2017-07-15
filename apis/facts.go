package apis

import (
	"strconv"

	"../app"
	"../models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	factService interface {
		Get(rs app.RequestScope, id int) (*models.Fact, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Fact, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Fact) (*models.Fact, error)
		Update(rs app.RequestScope, id int, model *models.Fact) (*models.Fact, error)
		Delete(rs app.RequestScope, id int) (*models.Fact, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	factResource struct {
		service factService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeFactResource(rg *routing.RouteGroup, service factService) {
	r := &factResource{service}
	rg.Get("/facts/<id>", r.get)
	rg.Get("/facts", r.query)
	rg.Post("/facts", r.create)
	rg.Put("/facts/<id>", r.update)
	rg.Delete("/facts/<id>", r.delete)
}

func (r *factResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *factResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *factResource) create(c *routing.Context) error {
	var model models.Fact
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *factResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *factResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
