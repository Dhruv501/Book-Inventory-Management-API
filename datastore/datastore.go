package datastore

import (
	"github.com/Dhruv501/Zopsmart-mini-Project/model"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type bookstore struct {
	db *gorm.DB
}

func NewBookStore() Book {
	dsn := "root:root123@tcp(localhost:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.Book{})

	return &bookstore{db: db}
}

func (b *bookstore) GetByID(ctx *gofr.Context, id uint) (*model.Book, error) {
	var book model.Book
	result := b.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (b *bookstore) Create(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	result := b.db.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (b *bookstore) Update(ctx *gofr.Context, book *model.Book) (*model.Book, error) {
	result := b.db.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (b *bookstore) Delete(ctx *gofr.Context, id uint) error {
	result := b.db.Delete(&model.Book{}, id)
	return result.Error
}
func (b *bookstore) GetAll(ctx *gofr.Context) ([]*model.Book, error) {
	var books []*model.Book
	if err := b.db.Select("*").Find(&books).Error; err != nil {
		ctx.Logger.Errorf("Error fetching books from database: %v", err)
		return nil, errors.DB{Err: err}

	}
	return books, nil
}
