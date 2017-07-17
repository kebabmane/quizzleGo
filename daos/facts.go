package daos

import (
	"../app"
	"../models"
)

// ArtistDAO persists artist data in database
type FactDAO struct{}

// NewArtistDAO creates a new ArtistDAO
func NewFactDAO() *FactDAO {
	return &FactDAO{}
}

// Get reads the artist with the specified ID from the database.
func (dao *FactDAO) Get(rs app.RequestScope, id int) (*models.Fact, error) {
	var fact models.Fact
	err := rs.Tx().Select().Model(id, &fact)
	return &fact, err
}

// Create saves a new artist record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *FactDAO) Create(rs app.RequestScope, fact *models.Fact) error {
	fact.Id = 0
	return rs.Tx().Model(fact).Insert()
}

// Update saves the changes to an artist in the database.
func (dao *FactDAO) Update(rs app.RequestScope, id int, fact *models.Fact) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	fact.Id = id
	return rs.Tx().Model(fact).Exclude("Id").Update()
}

// Delete deletes an artist with the specified ID from the database.
func (dao *FactDAO) Delete(rs app.RequestScope, id int) error {
	fact, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(fact).Delete()
}

// Count returns the number of the artist records in the database.
func (dao *FactDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("facts").Row(&count)
	return count, err
}

// Query retrieves the artist records with the specified offset and limit from the database.
func (dao *FactDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Fact, error) {
	facts := []models.Fact{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&facts)
	return facts, err
}
