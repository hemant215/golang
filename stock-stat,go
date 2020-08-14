
package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
   
    response, err := http.Get("https://query1.finance.yahoo.com/v7/finance/download/V?period1=1534204800&period2=1597363200&interval=1d&events=history")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

   
    _ , error := io.Copy(os.Stdout, response.Body)
    if error != nil {
        log.Fatal(error)
    }

    
}
