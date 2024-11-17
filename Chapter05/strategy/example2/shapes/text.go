package shapes

import (
	"log"

	strategy "github.com/ibiscum/Go-Design-Patterns/Chapter05/strategy/example2"
)

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	_, err := t.Writer.Write([]byte("Circle"))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
