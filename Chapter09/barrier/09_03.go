package barrier

// "fmt"
// "io/ioutil"
// "net/http"
// "time"

// var timeoutMilliseconds int = 5000

// var hasError bool
// for i := 0; i < requestNumber; i++ {
// 	resp := <-in
// 	if resp.Err != nil {
// 		fmt.Println("ERROR: ", resp.Err)
// 		hasError = true
// 		break
// 	}
// 	responses[i] = resp
// }

// if !hasError {
// 	for _, resp := range responses {
// 		fmt.Println(resp.Resp)
// 	}
// }
// }

// func makeRequest(out chan<- barrierResp, url string) {
// 	res := barrierResp{}
// 	client := http.Client{
// 		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
// 	}

// 	resp, err := client.Get(url)
// 	if err != nil {
// 		res.Err = err
// 		out <- res
// 		return
// 	}

// 	byt, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		res.Err = err
// 		out <- res
// 		return
// 	}

// 	res.Resp = string(byt)
// 	out <- res
// }
