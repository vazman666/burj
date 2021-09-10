package pkg

import (
	"burj/models"

	"github.com/tealeg/xlsx/v3"
)

func CreateXLSX() {
	handle := models.Tiss{
		Firm:       "производитель",
		PartNumber: "Партнам",
		Name:       "Наименование",
		Price:      "Цена",
		Qty:        "Количество",
		Amount:     "сумма",
		Remark:     "Кому",
	}

	wb := xlsx.NewFile() //создаём новый экскиз экселя

	sheetTest, err := wb.AddSheet("Sheet") //добавляем страничку
	if err != nil {
		panic(err)
	}
	sheetTest.SetColWidth(0, 5, 12.5)

	row1 := sheetTest.AddRow()
	_ = row1.WriteStruct(&handle, -1)
	row1.SetHeight(15)
	for i := 0; i < 6; i++ {
		_ = sheetTest.SetColAutoWidth(i, xlsx.DefaultAutoWidth)
	}
	for _, value := range models.Price {

		row1 = sheetTest.AddRow()        //добавляем строку
		_ = row1.WriteStruct(&value, -1) //и вставляе в эту строку строку из прайс
		row1.SetHeight(15)

	}
	err = wb.Save("origin.xlsx")
	if err != nil {
		panic(err)
	}

}
