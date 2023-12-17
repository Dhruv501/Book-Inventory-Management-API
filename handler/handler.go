package handler

import (
	"fmt"
	"strconv"

	"github.com/Dhruv501/Zopsmart-mini-Project/datastore"
	"github.com/Dhruv501/Zopsmart-mini-Project/model"

	"gofr.dev/pkg/gofr"
)

type bookHandler struct {
	store datastore.Book
}

func NewBookHandler(s datastore.Book) bookHandler {
	return bookHandler{store: s}
}

func (h bookHandler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	bookID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid ID: %s", id)
	}

	book, err := h.store.GetByID(ctx, uint(bookID))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (h bookHandler) Create(ctx *gofr.Context) (interface{}, error) {
	var book model.Book
	if err := ctx.Bind(&book); err != nil {
		return nil, fmt.Errorf("Invalid parameter: body")
	}

	createdBook, err := h.store.Create(ctx, &book)
	if err != nil {
		return nil, err
	}

	return createdBook, nil
}

func (h bookHandler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	bookID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid ID: %s", id)
	}

	var book model.Book
	if err := ctx.Bind(&book); err != nil {
		return nil, fmt.Errorf("Invalid parameter: body")
	}

	book.ID = uint(bookID)
	updatedBook, err := h.store.Update(ctx, &book)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (h bookHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	bookID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid ID: %s", id)
	}

	err = h.store.Delete(ctx, uint(bookID))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (h bookHandler) GetAll(ctx *gofr.Context) (interface{}, error) {
	books, err := h.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
