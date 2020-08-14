//script is to download 2 years visa stock data from yahoo.com
package main

// import required packages
import (
    "fmt"
    "net/http"
    "bytes"
	"io/ioutil"
)

func main() {

    //pass web link to read data
    response, _ := http.Get("https://query1.finance.yahoo.com/v7/finance/download/V?period1=1534204800&period2=1597363200&interval=1d&events=history")

    // first need to convert resposne data to byte else you will get conversion error while saving to csv
    buffer := new(bytes.Buffer)
    buffer.ReadFrom(response.Body)
    newStr := buffer.String()

	stockdata := []byte(newStr)
	
	err := ioutil.WriteFile("visastocks.csv", stockdata, 0777)    
	if err != nil {
  // print it out
  fmt.Println(err)
  
      }
	 
	
}
