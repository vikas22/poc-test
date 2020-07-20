package common

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"

)

// Repo - Repository containing methods for persistence layer
type Repo struct {
	Db *gorm.DB
}

// FindByPaymentID loads into model by given paymentID and service.
func (r *Repo) FindByPaymentID(ctx context.Context, m interface{}, id string, modelname string) error {

	q := r.Db.Model(m).Where(map[string]interface{}{
		"payment_id": id,
	}).Order("created_at DESC").First(m)

	if q.RecordNotFound() {
		err := fmt.Errorf("no records found")

		return err
	}
	return nil
}

// FetchLastEntity loads into model by given paymentID and service.
func (r *Repo) FetchLastEntity(ctx context.Context, m interface{}, id string, modelname string) error {

	q := r.Db.Model(m).Where(map[string]interface{}{
		"payment_id": id,
	}).Order("id DESC").First(m)

	if q.RecordNotFound() {
		err := fmt.Errorf("no records found")

		return err
	}
	return nil
}

// FindByID loads into model by given id and service.
func (r *Repo) FindByCondition(ctx context.Context, m interface{}, condition map[string]interface{}, modelname string) error {

	q := r.Db.Model(m).Where(condition).First(m)

	if q.RecordNotFound() {
		err := fmt.Errorf("no records found")
		return err
	}
	return nil
}

// Create - create the model in database
func (r *Repo) Create( m interface{}) error {

	if err := r.Db.Model(m).Create(m).Error; err != nil {
		return err
	}
	return nil
}

