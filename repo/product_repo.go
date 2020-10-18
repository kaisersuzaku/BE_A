package repo

import (
	"context"

	"github.com/kaisersuzaku/BE_A/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo struct {
	db *gorm.DB
}

func BuildProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		db: db,
	}
}

type IProductRepo interface {
	Read(ctx context.Context, id uint, product *models.Product)
	Begin(ctx context.Context) *gorm.DB
	Rollback(ctx context.Context, tx *gorm.DB) error
	Commit(ctx context.Context, tx *gorm.DB) error
	ReadForUpdateByID(ctx context.Context, tx *gorm.DB, id uint, product *models.Product) error
	Update(ctx context.Context, tx *gorm.DB, product models.Product) error
}

func (p ProductRepo) Begin(ctx context.Context) *gorm.DB {
	return p.db.Debug().Begin()
}

func (p ProductRepo) Rollback(ctx context.Context, tx *gorm.DB) error {
	return tx.Debug().Rollback().Error
}

func (p ProductRepo) Commit(ctx context.Context, tx *gorm.DB) error {
	return tx.Debug().Commit().Error
}

func (p ProductRepo) ReadForUpdateByID(ctx context.Context, tx *gorm.DB, id uint, product *models.Product) error {
	return tx.Debug().Model(models.Product{}).Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).Find(product).Error
}

func (p ProductRepo) Update(ctx context.Context, tx *gorm.DB, product models.Product) error {
	return tx.Debug().Model(models.Product{}).Where("id = ? AND products.stock > ?", product.ID, 0).Updates(product).Error
}

func (p ProductRepo) Read(ctx context.Context, id uint, product *models.Product) {
	p.db.Debug().Where(&models.Product{ID: id}).First(product)
}
