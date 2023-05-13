package main

import (
	"customManageSys/view"
)

func main() {
	customerView := view.NewCustomView()
	customerView.MainMenu()
}
