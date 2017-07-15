package services

import (
	"../app"
	"../models"
)

// artistDAO specifies the interface of the artist DAO needed by ArtistService.
type factDAO interface {
	// Get returns the artist with the specified artist ID.
	Get(rs app.RequestScope, id int) (*models.Fact, error)
	// Count returns the number of artists.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Fact, error)
	// Create saves a new artist in the storage.
	Create(rs app.RequestScope, artist *models.Fact) error
	// Update updates the artist with given ID in the storage.
	Update(rs app.RequestScope, id int, artist *models.Fact) error
	// Delete removes the artist with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// ArtistService provides services related with artists.
type FactService struct {
	dao factDAO
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewFactService(dao factDAO) *FactService {
	return &FactService{dao}
}

// Get returns the artist with the specified the artist ID.
func (s *FactService) Get(rs app.RequestScope, id int) (*models.Fact, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new artist.
func (s *FactService) Create(rs app.RequestScope, model *models.Fact) (*models.Fact, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the artist with the specified ID.
func (s *FactService) Update(rs app.RequestScope, id int, model *models.Fact) (*models.Fact, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the artist with the specified ID.
func (s *FactService) Delete(rs app.RequestScope, id int) (*models.Fact, error) {
	artist, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return artist, err
}

// Count returns the number of artists.
func (s *FactService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the artists with the specified offset and limit.
func (s *FactService) Query(rs app.RequestScope, offset, limit int) ([]models.Fact, error) {
	return s.dao.Query(rs, offset, limit)
}
