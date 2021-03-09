package datastruct

import "fmt"

// Pair implementation
type Pair struct {
	First  int
	Second int
}

func (p *Pair) String() string {
	return fmt.Sprintf("{%v, %v}", p.First, p.Second)
}
