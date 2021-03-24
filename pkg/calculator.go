package pkg

type Calculator interface {
	Add(a int, b int) int
}

type calculator struct{}

func (c calculator) Add(a int, b int) int {
	return a + b
}

func NewCalculator() Calculator {
	return &calculator{}
}
