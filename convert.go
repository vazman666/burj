package main

import (
	"burj/pkg"
)

func main() {
	pkg.LoadOrder()
	pkg.LoadNakl()
	pkg.Parse()
	pkg.CreateXLSX()

}
