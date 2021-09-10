package pkg

import (
	"burj/models"
	"fmt"
)

func Parse() {
	models.Price = make([]models.Tiss, 1)

	found := false

	for _, val := range models.Nakl {

		for _, val1 := range models.Order {
			if val1.PartNumber == val.PartNumber && val.PartNumber != "" {
				if val1.Price == val.Price {

					models.Price = append(models.Price, val1)
					found = true

				}
			}
			if !found {
				fmt.Errorf("НЕ НАЙДЕНО %v  в количистве %v\n", val.PartNumber, val.Qty)
			}
		}

	}

}
