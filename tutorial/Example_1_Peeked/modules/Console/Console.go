package Peeked

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func Handle(ch chan os.Signal) {
	signal.Notify(ch, os.Interrupt)
	for k := <-ch; ; k = <-ch {
		switch k {
		case os.Interrupt:
			fmt.Println("[+] Killing process and exiting")
			os.Exit(0)
		case os.Kill:
			fmt.Println("[*] Recived a signal to kill the program, exiting")
			os.Exit(0)
		}
	}
}

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Peeked> ")
		t, _ := reader.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		go Handle(make(chan os.Signal, 1))
		if strings.Compare(t[0:3], "set") == 0 {
			cmd := string(strings.TrimSpace(t[4:8]))
			val := string(strings.TrimSpace(t[8:]))
			ft := Func_map[cmd]
			ft.(func(string))(val)
		}
		if strings.Compare(t[0:7], "console") == 0 {
			cmd := string(strings.TrimSpace(t[7:]))
			ft := Console_Map[cmd]
			if ft != nil {
				c := ft.(func())
				c()
			}
		}
		if strings.Compare(t[0:6], "script") == 0 {
			ft := Script_Map[string(strings.TrimSpace(t[6:]))]
			if ft != nil {
				c := ft.(func())
				c()
			}
		}
	}
}
