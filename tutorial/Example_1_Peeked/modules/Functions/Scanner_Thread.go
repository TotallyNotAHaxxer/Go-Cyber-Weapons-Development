package Peeked

import (
	"fmt"
	XML "main/XML"
	"strconv"
)

func (Console_Import *Console_Data) Thread_Call() {
	pt := make(chan int, 65535)
	re := make(chan int)
	var open []int
	for o := 0; o < cap(pt); o++ {
		go Console_Import.Scan_Addr(pt, re)
	}
	go func() {
		for k := 1; k <= 65535; k++ {
			pt <- k
		}
	}()
	b := XML.Opener("XML/Database.xml")
	for p := 0; p < 65535; p++ {
		pts := <-re
		if pts != 0 {
			open = append(open, pts)
			port := strconv.Itoa(pts)
			XML.Find_and_Search("XML/Database.xml", port, b)
		}
	}
	fmt.Println("All open ports - ")
	for _, k := range open {
		fmt.Printf("\t[%v] \t %v\n", len(open), k)
	}
	fmt.Println("Total ports scanned -> ", scanned)

}
