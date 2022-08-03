package super_XML_TYPES

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Opener(filename string) []byte {
	f, x := os.Open(filename)
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	B, _ := ioutil.ReadAll(f)
	return B
}

func Find_and_Search(filename string, info string, b []byte) {
	var s Registry
	xml.Unmarshal(b, &s)
	fmt.Printf("Port - %s \n", info)
	for i := 0; i < len(s.Record); i++ {
		num := s.Record[i].Number
		if strings.Compare(info, num) == 0 {
			fmt.Printf("             | %s \t| %s   \t| %s\n", s.Record[i].Protocol, s.Record[i].Number, s.Record[i].Name)
		}
	}
	fmt.Printf("\n")

}
