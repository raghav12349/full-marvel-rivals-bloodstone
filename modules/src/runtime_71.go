package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.0" )

func ImsdQkdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGeTdnzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNqBx9ZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOwm65mUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CI8lUFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80LF4dQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHndGUObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAf87cuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0JPjWFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Z5vekNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBjokHXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ukss2BrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlUW0W7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYyIlwlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3xjoxyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9i2bbVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4XTav7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrPYs0D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4FyO3EZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SglFfRQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhDqm6d9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knp9y7iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKWMv4p1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCMC0SdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgiTbrsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqIuNmLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLkAFclFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igyf5iCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0NVVByeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yiLR26aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gnjmPykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SozdanGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wb7AknuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqWG7ohvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEKwkankWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4p2zaczLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0guPotGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7AyqCkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKFRufMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJs1tzByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7EZC2RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CvCKbqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54LDK9uKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQzgXU6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNN8faEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxrN5jceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMMAt9w9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFL7RffrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhkCoka8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nyywj5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func to6IOb9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHSCoHtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m52s86ZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQGdPqjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQo9DwikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hi9NPFwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnceteBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIgP2XLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6up6NbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQ5RPRZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6f5VBNteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihxHHdyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2a8tHdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTbLyAbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OX1OxpPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4NMRN9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYbFR6v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLye30mbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUHXlHRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAbT1OkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAbWmsBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jziyLkjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhK4CENbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cxt9zynXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlFXcnKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func II1kd7zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwO0ALEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oYYz9zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwBgC8D5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjBR29i3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bxa2DrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZ2GkDOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yC11HzV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gBLkMmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhiZkJjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Msjfp8C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DowT7JraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwSOrmVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QA4obstMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEu1OGgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsdYRrwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bvx9igdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNQzxDPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeRIjFeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyEjWoljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TphB6xUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zc03GqJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8fspU75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yr13YjGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLz1DD9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyUwFl2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBfSQbTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVaIgUSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ufi13ifBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EM9h3a1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALFJMyLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6OLbD1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpyG2euEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PciyG9fRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ifx1OyMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oNYZv1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNvK4RBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sowXcF1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3BLSzY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yCcSm5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtH43RRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n06G3QfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nElIZmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ow8x77uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5wtQvOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4875yWIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgDxe2hDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yn7SRwzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B1W0WzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDHjBcR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqpv6saQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Pz5NnDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkZzEB6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNQZznb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H5R2CeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhmFwBuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IUw1M38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2kF8mjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHYziTUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIphqoE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPvjwEUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLreVzzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NtjQZdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGZvHVULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wX611MEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCS9GRAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciLqAG2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaEbNmSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PtYNM5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OcYvOgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H15byJqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zxh6o5vDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTXuda9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AY9XcX3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrKwZg53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2IhxrM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6ChuoD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEQRAInKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJxJ3uqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46Eu4cWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2D6Nn28DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aoz0iuecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3N5HoSg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzlNzIn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvjz1jTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bt5UQQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJEFRjS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwyHpUmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Np6bjacKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLYoylszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuarwjgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bijJdlErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voq61b0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gR8VRH9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DscggYgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frkI24FfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOvrdQYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAL0KrpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOFcA9dYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDCh0NQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAgFL7XcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yArsUmgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIJMGzqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7WUI97WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMRxDWshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK4YOLRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8xWNZgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aw6nnz3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDKL2QM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDTIsMMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bz5Lh2gKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlZFHnaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xe2kqt41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZT73MrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxXOn3meWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FODGwsZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bz2DlDt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZLDomJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goOOXLD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IQxYe84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTAskfrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzNWw4kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQbMoG0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oo1aVDChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXSstGVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbAkCOBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ekjNygSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 181wufSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28PA5ujeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zxaIU70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hK3VeP0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdBYFZ8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTXE5R9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hDlBPpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drPWWefnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpM6JbrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7CtEahbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPHl1zm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YM0afLxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KT34ow9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaMqlAwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWvQjTUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0lJl2I4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V80415mPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLPUpBEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yectaXSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YA6NGA5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PT8pwh4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UU779U6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTiBJ3R7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdUM1DhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtnCciA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7KcL5GJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAXRxpfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UqmGjJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0d6w8D8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 562plsE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72aXEQjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFJbNFUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8K6CqSzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mGMSMeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KM1c0rsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNvoEMBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZyYfPdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9rf7DU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFAOWjFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmqZGrruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrXX6mONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tL9t3iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6DrJYjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFyaZeOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pp4lQW7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TACPaaOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0ADia5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QE8Wn8OnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3aDtMR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyFMlRasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUOTuxS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX2RL7NwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIZGsyFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0dBn5mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BVRZsQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Algy2Sp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kpof4P7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwzTaDZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRgTtlUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhpFA1SZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeXjeYirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaQxAO3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xcZ4P5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4F6rJ3uDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PStr5ee8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYuQuiBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMdapcUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOKtiSbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVU0YGu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWZnRarTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVjDoZO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xwn1e6zBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPPGzSRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49LZVTbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsZkTEaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJrjFyl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2qFW7tDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8xIrxvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2NDh7u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6mlJiL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvQCUlo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiJttnJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwR6qnluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7yXJJDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJ0sbSkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo4e1rp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeGg1HHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ag2mIKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pb2Pen6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLlnO1BXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiRniPx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHMGFBXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfEqxmOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYYAQoWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhcEJJiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuSFG3ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUmkXrp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcwXQ6MwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOO9hDKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSCMaReBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3kelik3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31e1KnRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tBedukoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGFoQ1BGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKQ98u41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqZjmMn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5khvu2OJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C77dbTLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpofB0SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjxICpozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0dfyrWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFlMSUjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DA41PNwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSmGwlM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iyaey9FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUiqNLtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rK5BgyvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNCX2a8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xY2WcUN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98Na8zRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zwo7f3u8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlNW9GtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WT69QM6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XctUau5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JpVZKD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9zyumqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugQ43HaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vVNMd1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNILRfrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mC3Nwt8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7X2TE4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eV2FJHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOSv6NjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZHut2akWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwNJdX3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miwieB0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k44LvigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPt9NOpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DgOByvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrwBfyHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvthwDojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CeeoyC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqM6gjTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THBR3x6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WlvoVsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E33cd7xWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOwUUqkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujROMgpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8F3gy09lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaFVuqnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqHdHfNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKo8f56lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TE4bW8QrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5URkMZFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rziaV5ppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0Vm0m3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqxapMBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCZL2ocVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQEBVTjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AWSFV8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bi2FEVGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlE2pz1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cx4ZnV5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OVmGjnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRz2dIJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kElBCo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEwwOvU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egvRlFJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtCnd5gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Syi4FNOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBBeKWGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qql5r5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCuHDlVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAqVNa9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJILRynfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PNZ8PL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOPElBDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtW37GsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DC5Cz7ghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MElU2vJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Og7PbUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZTPz19TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8vvPtNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0DttBHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4tUrLjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqSqqqeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgNPe2hwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wK4E62RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOzQ0THxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl7gY0diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMFsXbPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFalmXOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cK6nTyH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func om4F62g3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObDZVgMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQSF6bnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7BC1ZipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrbC6svOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Yb0nT1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYGF2dp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wqgCgXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DNWAvulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjbRymPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuKsFUWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qj9vYTmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOtkTHfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H09aDO0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3MkfX0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sE6ni8WvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujo7QcvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkg9hyTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTEuBhpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZ9q7KOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19GawK2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxSYnFmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3nPDnimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vg79WLQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPejJtrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVhD7RzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVgY1zxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBuwT9k5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMibAVpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2LWAy0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2HuIaTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LAtyu0JCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OA80mDSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIp5NvUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sscUhoQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvHG2VUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieZBpY4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmuAqvGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQUcLBFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8EFIXGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjdbn2syWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmDgxqCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajxFAEO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DS4AzeqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UxISSm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKrpo6eoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efrEduJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvLZIRVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNHz25DMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftMmpsClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mILdoz8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXRMxJRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Dc2rzUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6iL86AeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFU48xlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkxNVK5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blAefkAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Gg8VOpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkJiAv1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2ht253VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9D967SdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCjwFezaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2ItOHjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIgPJgzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mkeq1JkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNJPYgo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6VYuSXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r95eYzaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05m1zqp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4N0U5QDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJ7ZyGNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHCpKR7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CAtRM3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZAar0K2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqaAxxIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJ4GIgiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoVu478CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pULOVrCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qm8EKnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaIyC1ydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYzYzhghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F34ypfEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KoAsXFVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TuSoMBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ry6wmUZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38YM1UpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOTSL6ykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0LWHHgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kGhvav3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Um9e7kYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McNDWO0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwgIVcCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFeoeh7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YVtTxPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajS6Ns4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8buqpry8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0n9n1vgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWUxPgyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wG18QEcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVh2PDEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1fIg3PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UavuZ6sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80fwQZYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Drepj7VXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FI6ThdoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUTjDDQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsH6VqSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98oOja54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dNfBusVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1mEFcnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCQgRPBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3Roq0NNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fe5au1Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NVtwvMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXS9Y5ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYIQ0B9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAepgxNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxEZ21CDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2JeIdTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOcwaKKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9BK2jmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqD8CrbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38lUJWaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oj6YVP7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqiMY6TcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMXlbGEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXqSMZjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W16Had9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLuFWb69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46AkzM7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQbUtpHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlC7402QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Jop4OluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZy2xPd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grsaLhfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func capK5MgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hqQotmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sl47mc0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 125uQu0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeQPWE8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCKcIm4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHAgT3zwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7Iu1C9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUjlvaHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JF6YMu1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daRSZ3xEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dfd8jjDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5ia9x5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXo50mQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xltlVy07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atUsuu1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzqSJe8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgKXhZYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

