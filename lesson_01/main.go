package main

import (
    "fmt"
    "log"
    "net/http"
	"strconv"

    "github.com/julienschmidt/httprouter"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num_str := p.ByName("num")
	num_int, _ := strconv.Atoi(num_str)

	result_num_int := num_int % 2

	result_value := strconv.Itoa(result_num_int)

	fmt.Fprintf(w, result_value)

	// switch {
	// case num % 3 == 0 && num % 5 == 0:
	// 	return "FizzBuzz!"
	// case num % 3 == 0:
	// 	return "Fizz"
	// case num % 5 == 0:
	// 	return "Buzz"
	// default:
	// 	return string(num)
	// }
}

func main() {
    router := httprouter.New() // HTTPルーターを初期化

    router.GET("/FizzBuzz/:num", FizzBuzz)

    // Webサーバーを8080ポートで立ち上げる
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }
}