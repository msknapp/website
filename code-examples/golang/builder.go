package golang

import "errors"

type Thing struct {
	text   string
	number int
	array  []float64
}

func (t *Thing) Text() string {
	return t.text
}

func (t *Thing) Number() int {
	return t.number
}

type ThingBuilder struct {
	text   string
	number int
	array  []float64
}

func StartSomething() *ThingBuilder {
	return &ThingBuilder{
		array: make([]float64, 3),
	}
}

func (b *ThingBuilder) Text(s string) *ThingBuilder {
	b.text = s
	return b
}

func (b *ThingBuilder) Number(x int) *ThingBuilder {
	b.number = x
	return b
}

func (b *ThingBuilder) AddValue(x float64) *ThingBuilder {
	b.array = append(b.array, x)
	return b
}

func (b *ThingBuilder) Build() (*Thing, error) {
	if b.text == "" {
		return nil, errors.New("text is not set")
	}
	if b.number < 1 {
		// example of setting a default.
		b.number = 1
	}
	// defensive copies of slices.
	c := make([]float64, len(b.array))
	copy(c, b.array)
	return &Thing{
		text:   b.text,
		number: b.number,
		array:  c,
	}, nil
}
