//package main

import (
	_ "bytes"
    _ "encoding/json"
    "fmt"
    _ "github.com/google/uuid"
    _ "github.com/gorilla/websocket"
    _ "log"
    _ "io/ioutil"
    _ "net/http"
    _ "os"
    "regexp"
    "strings"
    _ "sync"
    _ "time"
)

func DropFirstLine(paragraph string) string {
	newLine := regexp.MustCompile("\\n")
	lines := newLine.Split(paragraph, -1)
	rest := lines[1:]
	return strings.Join(rest, "\\n")
}

func Erase(regex, s string) string {
    r, err := regexp.Compile(regex);
    if err != nil {
		fmt.Printf("Failed to compile regex: %s, err: %v", regex, err)
		return ""
    }
    erased := r.ReplaceAll([]byte(s), nil)
    return string(erased)
}

func Capture(regex, s string) string {
	re, err := regexp.Compile(regex);
	if err != nil {
		fmt.Printf("Failed to compile regex: %s, err: %v", regex, err)
		return ""
	}
	captured := re.FindString(s)
	return captured
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Hi(name string) string {
    return "Hello, " + name
}

func Take(howmany int, slice []int) ([]int, []int) {
    firstPart := slice[0:howmany]
    rest := slice[howmany:]
    return firstPart, rest
}

func Partitions(size int, items []int) (partitions [][]int) {
    for len(items) > 0 {
        groupSize := Min(size, len(items))
        group, rest := Take(groupSize, items)
        fmt.Println("items: ", items)
        partitions = append(partitions, group)
        items = rest
    }
    return
}

func Printer(format string, args ...string) {
    fmt.Printf("My format string=%s and args=%", format, args)
}

