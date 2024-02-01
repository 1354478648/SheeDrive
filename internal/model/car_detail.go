package model

import "SheeDrive/internal/model/entity"

type CarDetailGetListInput struct {
	Page      int
	PageSize  int
	Year      string
	Brand     string
	Model     string
	Category  string
	LowPrice  *int64
	HighPrice *int64
}

type CarDetailGetListOutput struct {
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Total    int                `json:"total"`
	Items    []entity.CarDetail `json:"items"`
}
