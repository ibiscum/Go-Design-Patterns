package barrier

import (
	"bytes"
	"io"
	"log"
	"os"
)

func barrier(endpoints ...string) {
	//requestNumber := len(endpoints)

	//in := make(chan barrierResp, requestNumber)
	//defer close(in)

	//responses := make([]barrierResp, requestNumber)

	// for _, endpoint := range endpoints {
	// 	//go makeRequest(in, endpoint)

}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, reader)
		if err != nil {
			log.Fatal(err)
		}
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}
