package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.5" )

func Ok3bTLa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9ElH2CbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCbiohsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilyd0RTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8DDnO5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func selY5G66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5e4UefsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lowBKcpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcTUKZK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z00q9wnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0dXNReUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caAQKfFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4J7R4AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80LVrbfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcGxERNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHlGzkSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HW8Dm0FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFPRre8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uq1cwuicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PfCrR4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDcQJ08aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXfGWWE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dASWzj1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcjgHBTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqdqU2iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMdydt4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJy1r0WSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcxOT5P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zt6coitgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCgZjfhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3tsvzocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dn2sXCwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rivtqpbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJF9SxFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89oNjFJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qb2LAUzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puvGAtlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4lv1DT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GL9H7KksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjG0BARzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAaGkGCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTleci7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZfYmlQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIGtmwWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeVRwatLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCIJiMTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XahwTsk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtj5cLtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LjEpZfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNDBEinHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPbOCXn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCSLwgqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

