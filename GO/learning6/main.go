package main

import (
	"mypro/fileop"
	"mypro/getinput"
	"mypro/usetime"
)

func main() {
	fileop.TestFileRead()
	fileop.TestBufio()
	fileop.TestIoutil()
	fileop.TestFuncWrite()
	fileop.TestBufioWr()
	fileop.TestIoutilWR()

	getinput.GetInputByBufio()
	usetime.TestTime()
	usetime.TestTimeSub()

	fileop.TestCopyFile()
	fileop.CharactersCount()

	getinput.GetInput()
}
