package service

import (
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

type OrderBook struct {
	repo repository.IOrderBookRepository
}

func NewOrderBook(repo repository.IOrderBookRepository) *OrderBook {
	return &OrderBook{
		repo: repo,
	}
}

func (o OrderBook) GetOrderBooks() ([]model.OrderBook, error) {
	return o.repo.GetOrderBooks()
}

func (o OrderBook) GetOrderUserBooks() ([]model.UserOrderBooksResponse, error) {
	userBooks, err := o.repo.GetOrderUserBooks(false)
	if err != nil {
		return []model.UserOrderBooksResponse{}, err
	}

	res := combineBooks(userBooks, false)
	return res, nil
}

func (o OrderBook) GetOrderUserBooksLastMounth() ([]model.UserOrderBooksResponse, error) {
	userBooks, err := o.repo.GetOrderUserBooks(true)
	if err != nil {
		return []model.UserOrderBooksResponse{}, err
	}

	res := combineBooks(userBooks, true)
	return res, nil
}

// withReturned(bool) - с возвращенными книгами
func combineBooks(userBooks []model.UserOrderBooks, withReturned bool) []model.UserOrderBooksResponse {
	data := make(map[uint]*model.UserOrderBooksResponse)
	for _, ub := range userBooks {
		if _, ok := data[ub.ID]; !ok {
			data[ub.ID] = &model.UserOrderBooksResponse{
				ID:       ub.ID,
				Login:    ub.Login,
				FullName: ub.FullName,
				Book:     []model.BookWithDate{},
			}
		}
		book := model.BookWithDate{
			ID:          ub.Book.ID,
			Name:        ub.Book.Name,
			Description: ub.Book.Description,
			Author:      ub.Book.Author,
			OrderDate:   ub.Book.OrderDate,
			ReturnDate:  ub.Book.ReturnDate,
		}

		if withReturned {
			data[ub.ID].AddBook(book)
		} else if !ub.Book.ReturnDate.Valid {
			data[ub.ID].AddBook(book)
		}
	}

	result := make([]model.UserOrderBooksResponse, 0, len(data))
	for _, user := range data {
		result = append(result, *user)
	}
	return result
}
