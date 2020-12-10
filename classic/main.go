package main

import (
	"fmt"
    "os"
	"strings"
	"bufio"
)
func main() {
    if len(os.Args) == 1 {
        fmt.Println("\nERROR: Please add an argument.")
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
    var result string
    if newLine == true {
        args := strings.Split(s, "\\n")
        for _, v := range args {
            for i := 0; i < 8; i++ {
                for _, w := range v {
                    result += GetLine(1 + int(w-32)*9 + i)
                }
                fmt.Println(result)
                result = ""
            }
        }
    } else {
        for i := 0; i < 8; i++ {
            for _, v := range s {
                result += GetLine(1 + int(v-32)*9 + i)
            }
            fmt.Println(result)
            result = ""
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
