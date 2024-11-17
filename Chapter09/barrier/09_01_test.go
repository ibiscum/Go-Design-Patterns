package barrier

import (
	"fmt"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		fmt.Println(endpoints)

	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}
		fmt.Println(endpoints)

	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		fmt.Println(endpoints)

	})
}

func Test_captureBarrierOutput(t *testing.T) {
	type args struct {
		endpoints []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captureBarrierOutput(tt.args.endpoints...); got != tt.want {
				t.Errorf("captureBarrierOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
