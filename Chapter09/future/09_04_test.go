package future

import (
	"fmt"
	"testing"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}
	fmt.Println(future)

	t.Run("Success result", func(t *testing.T) {

	})

	t.Run("Error result", func(t *testing.T) {

	})
}
