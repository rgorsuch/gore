package main

import (
	_ "bytes"
    _ "encoding/json"
    "fmt"
	"time"

	"github.com/google/uuid"
    "golang.org/x/tools/go/packages"
    _ "github.com/gorilla/websocket"
	_ "github.com/go-test/deep"
    _ "log"
    _ "io/ioutil"
    _ "net/http"
    "net/url"
    _ "os"
    "regexp"
    "strings"
    _ "sync"
)

func init() {
    // ignore this unused package which is imported for convenience
    _ = uuid.New
    _ = url.URL{}
    _ = websocket.Conn{}
}


func isToday(t time.Time, location *time.Location) bool {
	var today, someday struct {
		year  int
		month time.Month
		day   int
	}

	today.year, today.month, today.day = time.Now().In(location).Date()
	someday.year, someday.month, someday.day = t.In(location).Date()

	return today == someday
}

func DropFirstLine(paragraph string) string {
	newLine := regexp.MustCompile("\\n")
	lines := newLine.Split(paragraph, -1)
	rest := lines[1:]
	return strings.Join(rest, "\\n")
}

// Erase the matching regex from the string
func Erase(regex, s string) string {
    r, err := regexp.Compile(regex);
    if err != nil {
		fmt.Printf("Failed to compile regex: %s, err: %v", regex, err)
		return ""
    }
    erased := r.ReplaceAll([]byte(s), nil)
    return string(erased)
}

// Capture the matching part of the string, or else ""
func Capture(regex, s string) string {
	re, err := regexp.Compile(regex);
	if err != nil {
		fmt.Printf("Failed to compile regex: %s, err: %v", regex, err)
		return ""
	}
	if captured := re.FindStringSubmatch(s); len(captured) > 1 {
		return captured[1]
	}
	return ""
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

