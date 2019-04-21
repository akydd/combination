package combination

import "testing"

func TestChoose(t *testing.T) {
	c := NewCombination()

	v := c.Choose(10, 2)
	if v != 45 {
		t.Errorf("Combination was incorrect, got %d, want %d.", v, 45)
	}
}

func TestChooseBigNumber(t *testing.T) {
	c := NewCombination()

	v := c.Choose(90, 18)
	if v != uint64(3789648142708598775) {
		t.Errorf("Choose(90, 18) is incorrect, got %v, want %v.", v, uint64(3789648142708598775))
	}
}

func TestCacheIsNotUpdatedForSameNandK(t *testing.T) {
	c := NewCombination()

	_ = c.Choose(10, 2)
	ts := c.GetUpdated()

	_ = c.Choose(10, 2)

	if ts != c.GetUpdated() {
		t.Error("The second call to Choose(10, 2) should not have updated the cache a second time.")
	}
}

func TestCacheIsUpdatedForSameNdifferentK(t *testing.T) {
	c := NewCombination()

	_ = c.Choose(10, 2)
	ts := c.GetUpdated()

	_ = c.Choose(10, 5)

	if ts == c.GetUpdated() {
		t.Error("A second, non repeated, call to Choose should have updated the cache.")
	}
}

func TestCacheIsUpdatedForDifferentNsameK(t *testing.T) {
	c := NewCombination()

	_ = c.Choose(10, 2)
	ts := c.GetUpdated()

	_ = c.Choose(12, 2)

	if ts == c.GetUpdated() {
		t.Error("A second, non repeated, call to Choose should have updated the cache.")
	}
}

func TestRepeatedCache(t *testing.T) {
	c := NewCombination()

	x := c.Choose(10, 2)
	y := c.Choose(12, 2)
	z := c.Choose(10, 2)

	if x != z {
		t.Errorf("Repeated calls to Choose(x, y) should return the same value. Wanted %d and %d, got %d and %d.", 45, 45, x, y)
	}
}

func TestIsNcached(t *testing.T) {
	c := NewCombination()

	_ = c.Choose(10, 2)

	if !c.IsNcached(10) {
		t.Error("n = 10 should have been cached.")
	}

	if c.IsNcached(2) {
		t.Error("n = 2 should not have been cached.")
	}
}

func TestIsCached(t *testing.T) {
	c := NewCombination()

	_ = c.Choose(10, 2)

	if !c.IsCached(10, 2) {
		t.Error("n=10 and k=2 should have been cached.")
	}
}
