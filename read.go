package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	sname := make([]Name, 0)
	fmt.Printf("Enter file name")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		fullName := strings.Fields(scan.Text())
		n := Name{fname: fullName[0], lname: fullName[1]}
		sname = append(sname, n)
	}
	for _, name := range sname {
		fmt.Printf("fname: %s; lname: %s\n", name.fname, name.lname)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

}
