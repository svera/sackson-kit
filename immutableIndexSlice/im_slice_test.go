package immutableIndexSlice

import (
	"testing"
)

func TestAppend(t *testing.T) {
	slice := New()
	slice.Append("item 1", "item 2", "item 3")
	if len(slice.container) != 3 {
		t.Errorf("Slice must have exactly 3 items, got %d", len(slice.container))
	}
}

func TestDelete(t *testing.T) {
	slice := New()
	slice.Append("item")
	slice.Delete(0)
	if len(slice.container) != 0 {
		t.Errorf("Slice must have no items, got %d", len(slice.container))
	}
}

func TestIndexesDoesntChangeAfterDelete(t *testing.T) {
	slice := New()
	slice.Append("item 1")
	slice.Append("item 2")
	slice.Append("item 3")
	slice.Delete(1)
	if len(slice.container) != 2 {
		t.Errorf("Slice must have exactly 2 items, got %d", len(slice.container))
	}
	if slice.Get(2).(string) != "item 3" {
		t.Errorf("Slice index 2 must be %s, got %s", "item 3", slice.Get(2).(string))
	}
}

func TestGet(t *testing.T) {
	slice := New()
	slice.Append("item")
	if slice.Get(0).(string) != "item" {
		t.Errorf("Slice.Get(0) must return 'item', got %s", slice.Get(0).(string))
	}
}

func TestLen(t *testing.T) {
	slice := New()
	slice.Append("item")
	if slice.Len() != 1 {
		t.Errorf("Slice.Len() must return 1, got %d", slice.Len())
	}
}

func TestGetMustPanicIfIndexDoesNotExist(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	slice := New()
	slice.Get(0)
}

func TestSortedKeysLoop(t *testing.T) {
	slice := New()
	expected := []int{0, 1, 2, 3, 4}
	slice.Append(0)
	slice.Append(1)
	slice.Append(2)
	slice.Append(3)
	slice.Append(4)

	i := 0
	for key := range slice.SortedKeys() {
		val := slice.Get(key).(int)
		if val != expected[i] {
			t.Errorf("Slice.Get(%d) must return %d, got %d", i, expected[i], val)
		}
		i++
	}
}
