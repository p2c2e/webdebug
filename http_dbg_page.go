package main

import (
    "time"
    "fmt"
    "net/http"
    "log"
    "net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "1.1.1.1:80") // 8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}

func debuginfo(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "<html><head/><body>")
    fmt.Fprintf(w, "%s<br/>", time.Now())
    fmt.Fprintf(w, "%s<br/>", GetOutboundIP())
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v<br/>", name, h)
        }
    }
    fmt.Fprintf(w, "</body></html>")
}

func main() {

    http.HandleFunc("/debug", debuginfo)
    http.ListenAndServe(":8888", nil)
}
