package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
)
// 有 cors 问题，看下 gin 怎么解决，跟着做下
func main() {

    http.HandleFunc("/flv/", playVideo)

    http.ListenAndServe(":80", nil)

}
func playVideo(w http.ResponseWriter, r *http.Request) {
     strs := strings.Split(r.URL.Path, "/")
     fmt.Println(strs)
     filePath :="./videos/" + strs[2]
     file,err := os.Open(filePath)
     if err != nil {
        log.Println(err)
        return
     }
     io.Copy(w,file)
}
