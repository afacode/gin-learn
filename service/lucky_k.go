package service

import (
	"gin-learn/global"
	"gin-learn/model"
	"gin-learn/model/response"
)

type Res struct {
	Date       []string  `json:"date"`
	Total      []float64 `json:"total"`
	Share      []float64 `json:"share"`
	Value      []float64 `json:"value"`
	Percentage []float64 `json:"percentage"`
}

func GetAllK() (err error, kMap response.LuckyKResponse) {
	var kList []model.LuckyK
	err = global.G_DB.Find(&kList).Error
	for _, v := range kList {
		kMap.Date = append(kMap.Date, v.Date)
		kMap.Total = append(kMap.Total, v.Total)
		kMap.Share = append(kMap.Share, v.Share)
		kMap.Value = append(kMap.Value, v.Value)
		kMap.Percentage = append(kMap.Percentage, v.Percentage)
	}
	return err, kMap
}
