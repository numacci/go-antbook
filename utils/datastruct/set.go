package datastruct

import (
	"strconv"
	"strings"
)

type Set struct {
	data map[interface{}]struct{}
}

func NewSet(cap int) *Set {
	return &Set{data: make(map[interface{}]struct{}, cap)}
}

func (s *Set) Insert(x interface{}) {
	s.data[x] = struct{}{}
}

func (s *Set) Erase(x interface{}) {
	delete(s.data, x)
}

func (s *Set) Contains(x interface{}) bool {
	_, ret := s.data[x]
	return ret
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) String() string {
	ret := make([]string, 0, len(s.data))
	for d := range s.data {
		switch d.(type) {
		case string:
			ret = append(ret, d.(string))
		case int:
			i := strconv.Itoa(d.(int))
			ret = append(ret, i)
		}
	}
	return "{" + strings.Join(ret, ", ") + "}"
}
