package models

//import "gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/golang.org/x/text@v0.3.6/number"
//const Providers=1\\количество обрабатываемых поставщиков
type Tiss struct {
	Firm       string
	PartNumber string
	Name       string //наименование
	Price      string //цена
	Qty        string //количество
	Amount     string //сумма
	Remark     string
}

var Price []Tiss
var Order []Tiss
