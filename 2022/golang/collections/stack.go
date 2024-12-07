package collections

import "errors"

type RuneStack struct {
	data *[]rune
}

func NewRuneStack() RuneStack {
	return RuneStack{&[]rune{}}
}

func (s RuneStack) Append(r rune) {
	*s.data = append(*s.data, r)
}

func (s RuneStack) Pop() (rune, error) {
	if len(*s.data) == 0 {
		return 0, errors.New("cannot pop from empty RuneStack")
	}
	r := (*s.data)[len(*s.data)-1]
	*s.data = (*s.data)[:len(*s.data)-1]
	return r, nil
}

func (s RuneStack) Size() int {
	return len(*s.data)
}

func (s RuneStack) Top() (rune, bool) {
	if len(*s.data) == 0 {
		return 0, false
	}
	return (*s.data)[len(*s.data)-1], true
}

func (s RuneStack) ToArray() []rune {
	return *s.data
}