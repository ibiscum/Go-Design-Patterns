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
