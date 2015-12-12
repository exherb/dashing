package main

import(
    "path/filepath"
    "flag"
    "fmt"
    "github.com/exherb/Dashing"
    _ "./jobs"
)

func main() {
    var host string
    flag.StringVar(&host, "host", "127.0.0.1", "bind host")
    var port int
    flag.IntVar(&port, "port", 8080, "bind port")
    flag.Parse()
    dirpath := flag.Arg(0)
    dirpath, _ = filepath.Abs(dirpath)

    err := dashing.Server(dirpath, host, port)
    if err != nil {
        fmt.Println(err)
    }
}
