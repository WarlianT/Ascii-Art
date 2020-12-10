package main

import (
	"fmt"
    "os"
	"strings"
	"bufio"
)
func main() {
    if len(os.Args) == 1 {
        fmt.Println("\nVeuillez entrer des caractères")
        return
    }
    s := os.Args[1]
    for _, v := range os.Args[2:] {
                s += " " + v
        }
    prev := '0'
    newLine := false
    for _, v := range s {
        if v == 'n' && prev == '\\' {
            newLine = true
        }
        prev = v
    }
    var resultat string
    if newLine == true {
        args := strings.Split(s, "\\n")
        for _, v := range args {
            for i := 0; i < 8; i++ {
                for _, x := range v {
                    resultat += GetLine(1 + int(x-' ')*9+i)
                }
                fmt.Println(resultat)
                resultat = ""
            }
        }
    } else {
        for i := 0; i < 8; i++ {
            for _, c := range s {
                resultat += GetLine(1 + int(c-' ')*9+i)
            }
            fmt.Println(resultat)
            resultat = ""
        }
    }
}
func GetLine(num int) string {
	f, e := os.Open("standard.txt")
	if e!= nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line
}
