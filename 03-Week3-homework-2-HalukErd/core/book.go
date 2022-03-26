package core

import (
	"fmt"
)

type BaseEntity struct {
	Id   int64
	Name string
}

type Author struct {
	BaseEntity
}

type Book struct {
	BaseEntity
	PageNumber int
	Stock      int
	Price      float64
	StockCode  int
	Isbn       string
	Author     Author
}

func (b Book) ToString() string {
	return fmt.Sprintf("Id:%d  Name:%s  Page Number:%d  Stock:%d  Price:%f  Stock Code:%d  ISBN:%s  Author:%s",
		b.Id,
		b.Name,
		b.PageNumber,
		b.Stock,
		b.Price,
		b.StockCode,
		b.Isbn,
		b.Author.Name,
	)
}

var Authors = []Author{
	{
		BaseEntity: BaseEntity{
			Id:   0,
			Name: "George RR Martin",
		},
	},
	{
		BaseEntity: BaseEntity{
			Id:   1,
			Name: "Patrick Rothfus",
		},
	},
	{
		BaseEntity: BaseEntity{
			Id:   2,
			Name: "Tolkien",
		},
	},
}

var Books = []Book{
	{
		BaseEntity: BaseEntity{
			Id:   0,
			Name: "A Song Of Ice And Fire",
		},
		PageNumber: 1200,
		Stock:      253,
		Price:      23.5,
		StockCode:  1513,
		Isbn:       "99921-58-10-7",
		Author:     Authors[0],
	},
	{
		BaseEntity: BaseEntity{
			Id:   1,
			Name: "The Kingkiller Chronicle",
		},
		PageNumber: 623,
		Stock:      123,
		Price:      54.2,
		StockCode:  12314,
		Isbn:       "9971-5-0210-0",
		Author:     Authors[1],
	},
	{
		BaseEntity: BaseEntity{
			Id:   2,
			Name: "The Lord Of The Rings",
		},
		PageNumber: 1366,
		Stock:      121,
		Price:      52.8,
		StockCode:  2131,
		Isbn:       "960-425-059-0",
		Author:     Authors[2],
	},
}
