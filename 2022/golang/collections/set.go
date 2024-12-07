package collections

type RuneSet struct {
	data map[rune]bool
}

func NewEmptyRuneSet() *RuneSet {
	return &RuneSet{map[rune]bool{}}
}

func NewRuneSetFromString(st string) *RuneSet {
	s := NewEmptyRuneSet()
	for _, r := range st {
		s.Add(r)
	}
	return s
}

func (s *RuneSet) Add(r rune) {
	s.data[r] = true
}

func (s *RuneSet) Remove(r rune) {
	delete(s.data, r)
}

func (s RuneSet) Contains(r rune) bool {
	return s.data[r]
}

func (s RuneSet) Union(args... RuneSet) *RuneSet {
	union := NewEmptyRuneSet()
	for k := range s.data {
		union.Add(k)
	}
	for _, arg := range args {
		for k := range arg.data {
			union.Add(k)
		}
	}
	return union
}

func (s RuneSet) Intersection(args... RuneSet) *RuneSet {
	inter := NewEmptyRuneSet()
	for k := range s.data {
		flag := true
		for _, arg := range args {
			if !arg.Contains(k) {
				flag = false
				break
			}
		}
		if flag {
			inter.Add(k)
		}
	}
	return inter
}

func (s RuneSet) Size() int {
	return len(s.data)
}

func (s RuneSet) Array() []rune {
	keys := make([]rune, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func StringIntersection(strs... string) *RuneSet {
	var s *RuneSet
	for i, str := range strs {
		if i == 0 { 
			s = NewRuneSetFromString(str) 
		} else {
			s = s.Intersection(*NewRuneSetFromString(str))
		}
	}
	return s
}
