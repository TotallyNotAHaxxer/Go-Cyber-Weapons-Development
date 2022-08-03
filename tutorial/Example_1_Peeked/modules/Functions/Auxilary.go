package Peeked

import (
	"fmt"
)

var C Console_Data

func (Console_Import *Console_Data) Set_Aux_Host(host string) {
	Console_Import.Host = host
	fmt.Println("\n[+] Setting host -> ", host)
}

func (Console_Import *Console_Data) Set_Aux_Verbosity(verbose string) {
	Console_Import.Verbose = verbose
	fmt.Println("\n[+] Setting Verbose -> ", verbose)
}

func (Console_Import *Console_Data) Settings() {
	fmt.Println("[--------- Config: Hostname -> ", Console_Import.Host)
	fmt.Println("[--------- Config: Verbose  -> ", Console_Import.Verbose)
}

func (Console_Import *Console_Data) Clear() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
}
