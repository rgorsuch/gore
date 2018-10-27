import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
    "github.com/gorilla/websocket"
    "log"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "strings"
    "sync"
    "time"
    "reno-hammer-client/structs"
)


func IgnoreThis() {
    _ = uuid.New()
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

