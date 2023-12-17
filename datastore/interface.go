package datastore

import (
	"github.com/Dhruv501/Zopsmart-mini-Project/model"

	"gofr.dev/pkg/gofr"
)

type Book interface {
	GetByID(ctx *gofr.Context, id uint) (*model.Book, error)
	Create(ctx *gofr.Context, book *model.Book) (*model.Book, error)
	Update(ctx *gofr.Context, book *model.Book) (*model.Book, error)
	Delete(ctx *gofr.Context, id uint) error
	GetAll(ctx *gofr.Context) ([]*model.Book, error)
}
