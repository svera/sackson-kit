package shufflable

import "testing"

func TestNew(t *testing.T) {
	shuff := New("item 1", "item 2")
	if len(shuff.Items) != 2 {
		t.Errorf("Shuff must have exactly 3 items, got %d", len(shuff.Items))
	}
}

func TestAppend(t *testing.T) {
	shuff := New("item 1", "item 2")
	shuff.Append("item 3")
	if len(shuff.Items) != 3 {
		t.Errorf("Shuff must have exactly 3 items, got %d", len(shuff.Items))
	}
}

func TestDelete(t *testing.T) {
	shuff := New("item 1", "item 2")
	shuff.Delete(1)
	if len(shuff.Items) != 1 {
		t.Errorf("Shuff must have exactly 1 item, got %d", len(shuff.Items))
	}
}

func TestDrawRemovesItem(t *testing.T) {
	shuff := New()
	for i := 0; i < 100; i++ {
		shuff.Append(i)
	}
	shuff.Draw()
	if len(shuff.Items) == 100 {
		t.Errorf("Shuff must have exactly 99 items after draw, got %d", len(shuff.Items))
	}
}

func TestDrawIsRandom(t *testing.T) {
	shuff := New()
	for i := 0; i < 1000; i++ {
		shuff.Append(i)
	}
	value1 := shuff.Draw()
	shuff = New()
	for i := 0; i < 1000; i++ {
		shuff.Append(i)
	}
	value2 := shuff.Draw()
	if value1 == value2 {
		t.Errorf("Randomly extracted values must be different [NOTE: This test could bring false positives]")
	}
}

func TestDrawEmpty(t *testing.T) {
	shuff := New()
	item := shuff.Draw()
	if item != nil {
		t.Errorf("Extracting a value from an empty set must return nil")
	}
}
