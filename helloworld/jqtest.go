package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/savaki/jq"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// key, _ := reader.ReadString('\n')
	key := os.Args[1]
	op, _ := jq.Parse(key)
	data := []byte(`{"hello":"world"}`)
	value, err := op.Apply(data)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println(string(value))
	}
	//array := new([][]int)
	// go func() {
	// 	for {
	// 		resp, err := http.Get("http://baidu.com")
	// 		if err != nil {
	// 			fmt.Print(err.Error())
	// 			break
	// 		} else {
	// 			body, _ := ioutil.ReadAll(resp.Body)
	// 			fmt.Println(string(body))
	// 			resp.Body.Close()
	// 		}
	// 		a := make([]int, 1024*1024)
	// 		*array = append(*array, a)
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println(len(*array))
	// 	}
	// }()
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		resp.Body.Close()
	}
	fmt.Println(os.Args[0])
	http.ListenAndServe("localhost:8502", nil)

}
