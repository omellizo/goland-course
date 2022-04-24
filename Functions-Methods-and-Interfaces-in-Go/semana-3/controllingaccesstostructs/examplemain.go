package main

import "controllingaccesstostructs"

func main() {
	var p controllingaccesstostructs.Point
	p.InitMe(5, 3)
	p.Scale(4)
	p.PrintMe()

}
