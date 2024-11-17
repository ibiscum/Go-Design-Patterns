package main

import (
	"log"
	_ "net/http/pprof"
	// "strings"
)

type PreffixSuffixWorker struct {
}

type Request struct {
	Data string
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	log.Printf("in: %#v", in)
	out := make(chan Request)
	go func() {
		// for msg := range in {
		// 	s, ok := msg.Data.(string)

		// 	if !ok {
		// 		msg.handler(nil)
		// 		continue
		// 	}

		// 	msg.Data = strings.ToUpper(s)

		// 	out <- msg
		// }
		// close(out)
	}()
	return out
}

func main() {
	in := make(chan Request)
	var w PreffixSuffixWorker
	out := w.uppercase(in)
	log.Printf("out: %#v", out)
}
