// Package shufflable contains the implementation of Shufflable, which holds
// a set of items and provides a Draw() method that returns a randomly picked item from that set,
// as well as the convenience methods Delete() and Append() to manage it.
package shufflable

import (
	"math/rand"
	"time"
)

// Shufflable stores items and provides methods to access and manage them.
type Shufflable struct {
	Items []interface{}
}

var rn *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rn = rand.New(source)
}

// New initialises and returns a Shufflable instance.
func New(its []interface{}) *Shufflable {
	return &Shufflable{Items: its}
}

// Draw extracts a random item from the set and returns it.
func (s *Shufflable) Draw() interface{} {
	remainingItems := len(s.Items)
	if remainingItems == 0 {
		return nil
	}

	pos := 0
	if remainingItems > 1 {
		pos = rn.Intn(remainingItems - 1)
	}

	item := s.Items[pos]
	s.Delete(pos)
	return item
}

// Delete removes passed item from the set.
func (s *Shufflable) Delete(i int) *Shufflable {
	s.Items = append(s.Items[:i], s.Items[i+1:]...)
	return s
}

// Append adds the passed items to the set.
func (s *Shufflable) Append(items []interface{}) *Shufflable {
	s.Items = append(s.Items, items...)
	return s
}
