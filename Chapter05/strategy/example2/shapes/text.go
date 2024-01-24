package shapes

import strategy "github.com/ibiscum/Go-Design-Patterns/Chapter05/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
