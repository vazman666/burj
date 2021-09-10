package pkg

import (
	"burj/models"
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx/v3"
)

func LoadNakl() {
	models.Nakl = make([]models.Tiss, 0)
	wb, err := xlsx.OpenFile("nakl.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["invoice"]
	if !ok {
		fmt.Println("Sheet Invoice does not exist")
		return
	}
	fmt.Printf("MaxRow in nakl =%v\n", sh.MaxRow)

	for i := 1; i < sh.MaxRow; i++ {
		//for i, val := range sh {
		//fmt.Printf("%v\n", i)
		theCell, err := sh.Cell(i, 2) //firm
		if err != nil {

			panic(err)
		}
		firm, err := theCell.FormattedValue()

		if err != nil {
			fmt.Printf("Ошибка преобразования firm строка %v\n", i)

			panic(err)
		}
		theCell, err = sh.Cell(i, 2) //partnumber
		if err != nil {
			continue
			panic(err)
		}
		partnumber, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 7) //remark
		if err != nil {
			continue
			panic(err)
		}
		remark, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 5) //name
		if err != nil {
			continue
			panic(err)
		}
		name, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 5) //qty
		if err != nil {
			panic(err)
		}
		qty, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}
		theCell, err = sh.Cell(i, 6) //price
		if err != nil {
			continue
			panic(err)
		}
		price, err := theCell.FormattedValue()
		if err != nil {
			continue
			panic(err)
		}

		qtyFloat, err := strconv.ParseFloat(qty, 2)
		if err != nil {

			continue
			//panic(err)
		}
		priceFloat, err := strconv.ParseFloat(price, 2)
		if err != nil {
			continue
			panic(err)
		}
		priceFloat = priceFloat * models.Kurs
		price = fmt.Sprintf("%5.2f", priceFloat)

		amount := fmt.Sprintf("%5.2f", qtyFloat*priceFloat)

		tmp := models.Tiss{Firm: firm, PartNumber: partnumber, Name: name, Price: price, Qty: qty, Amount: amount, Remark: remark}

		models.Nakl = append(models.Nakl, tmp)

	}
	fmt.Printf("Loads ok, Выкидываем дубликаты\n")

	for i, val := range models.Nakl {

		for j := i; j < len(models.Nakl); j++ {
			if i != j && val.Price == models.Nakl[j].Price && val.PartNumber == models.Nakl[j].PartNumber && val.Remark == models.Nakl[j].Remark {
				fmt.Printf("qty [%v]=%v qty[%v]=%v PartNum=%v\n", i, models.Nakl[i].Qty, j, models.Nakl[j].Qty, val.PartNumber)
				qtyi, _ := strconv.Atoi(val.Qty)
				qtyj, _ := strconv.Atoi(models.Nakl[j].Qty)
				models.Nakl[j].Qty = strconv.Itoa(qtyi + qtyj)
				fmt.Printf("j=%v  Qty[j]=%v\n", j, models.Nakl[j].Qty)
				models.Nakl = models.Nakl[:i+copy(models.Nakl[i:], models.Nakl[i+1:])] //удаляем j элемент

			}
		}
	}

	//fmt.Printf("Nakl = %v\n", models.Nakl)
}
