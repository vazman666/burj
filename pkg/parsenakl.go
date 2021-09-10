package pkg

import (
	"burj/models"
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

func Parse() {
	models.Price = make([]models.Tiss, 1)
	wb, err := xlsx.OpenFile("nakl.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["invoice"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Printf("MaxRow=%v\n", sh.MaxRow)
	count := 0
	for i := 1; i < sh.MaxRow; i++ {
		//for i, val := range sh {
		fmt.Printf("%v\n", i)
		theCell, err := sh.Cell(i, 2)
		if err != nil {
			panic(err)
		}
		partnumber, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 4)
		if err != nil {
			panic(err)
		}
		qty, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 5)
		if err != nil {
			panic(err)
		}
		price, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		for _, val := range models.Order {
			if val.PartNumber == partnumber && partnumber != "" {
				fmt.Printf("  Найдена строка %v\n\n", val)
				count++
			}
		}
		//fmt.Printf("i = %v  qty = %v\n", i, qty)
		fmt.Printf("Получилось строк %v qty=%v  price=%v\n", count, qty, price)
		/*qtyFloat, err := strconv.ParseFloat(qty, 2)
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
		models.Order = append(models.Order, tmp)*/
	}
	//fmt.Printf("Order = %v\n", models.Order)
}
