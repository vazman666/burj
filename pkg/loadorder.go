package pkg

import (
	"burj/models"
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx/v3"
)

func LoadOrder() {
	models.Order = make([]models.Tiss, 0)
	wb, err := xlsx.OpenFile("Orders.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["Orders"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Printf("MaxRow=%v\n", sh.MaxRow)

	for i := 1; i < sh.MaxRow; i++ {
		//for i, val := range sh {
		//fmt.Printf("%v\n", i)
		theCell, err := sh.Cell(i, 2)
		if err != nil {
			panic(err)
		}
		firm, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 3)
		if err != nil {
			panic(err)
		}
		partnumber, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 9)
		if err != nil {
			panic(err)
		}
		remark, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 16)
		if err != nil {
			panic(err)
		}
		name, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 20)
		if err != nil {
			panic(err)
		}
		qty, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 22)
		if err != nil {
			panic(err)
		}
		price, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("i = %v  qty = %v\n", i, qty)

		qtyFloat, err := strconv.ParseFloat(qty, 2)
		if err != nil {
			//fmt.Printf("Ошибка преобразования qty\n")
			continue
			//panic(err)
		}
		priceFloat, err := strconv.ParseFloat(price, 2)
		if err != nil {
			panic(err)
		}
		amount := fmt.Sprintf("%5.2f", qtyFloat*priceFloat)
		//sales165Str := strconv.Itoa(sales165 * 1.65)		//sales205Str := strconv.Itoa(sales205 * 2.05)
		tmp := models.Tiss{Firm: firm, PartNumber: partnumber, Name: name, Price: price, Qty: qty, Amount: amount, Remark: remark}

		models.Order = append(models.Order, tmp)

	}
	fmt.Printf("Order = %v\n", models.Order)
}
