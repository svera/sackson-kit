// Package immutableIndexSlice contains the implementation of ImSlice, which acts
// like a regular slice, with the exception that its indexes are immutable.
package immutableIndexSlice

import "sort"

// ImSlice stores items and provides methods to access and manage them.
type ImSlice struct {
	container map[int]interface{}
	counter   int
}

// New returns a new instance of ImSlice
func New() *ImSlice {
	return &ImSlice{
		container: map[int]interface{}{},
		counter:   0,
	}
}

// SortedKeys returns a slice with the keys sorted from ascendently.
// This method must be used when iterating over ImSlice contents this way:
//
//      for key := range slice.SortedKeys() {
//		    val := slice.Get(key).(int)  // Change the type assertion to whatever the slice is storing
//          ...
//      }
func (s *ImSlice) SortedKeys() []int {
	keys := make([]int, len(s.container))
	i := 0
	for k := range s.container {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

// Append adds a new item to the end of ImSlice
func (s *ImSlice) Append(item interface{}) *ImSlice {
	s.container[s.counter] = item
	s.counter++
	return s
}

// Delete removes the item at the passed position in ImSlice
func (s *ImSlice) Delete(index int) *ImSlice {
	delete(s.container, index)
	return s
}

// Len returns how many items imSlice is holding
func (s *ImSlice) Len() int {
	return len(s.container)
}

// Get returns the item at the passed position in ImSlice
func (s *ImSlice) Get(index int) interface{} {
	if val, ok := s.container[index]; ok {
		return val
	}
	panic("Index out of range")
}
