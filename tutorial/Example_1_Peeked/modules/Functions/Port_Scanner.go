package Peeked

import (
	"fmt"
	"net"
	"strconv"
)

// port scanner for the host
var scanned int

func (Console_Import *Console_Data) Scan_Addr(p, r chan int) {
	scanned = 0
	for l := range p {
		scanned++
		v := strconv.Itoa(l)
		conn, err := net.Dial("tcp", net.JoinHostPort(Console_Import.Host, v))
		if err != nil {
			if Console_Import.Verbose == "true" {
				fmt.Println(err)
			}
			r <- 0
			continue
		}
		conn.Close()
		r <- l
	}
}
