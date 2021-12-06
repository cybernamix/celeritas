package cache

import "testing"

func TestBadgerCache_Has(t *testing.T) {
	err := testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("Found 'foo' in cache when it should not be")
	}

	_ = testBadgerCache.Set("foo", "bar")
	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not in cache when it should be")
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}
}

func TestBadgerCache_Get(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testBadgerCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("did not get correct value from cache")
	}
}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo in cache when it should not be there")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("alpha", "blah")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("value in cache when it should be empty")
	}

}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	err := testBadgerCache.Set("gamma", "delta")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Set("gamma2", "omega")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("who", "me")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyByMatch("gam")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("gamma")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("gamma in cache when it should not be")
	}

	inCache, err = testBadgerCache.Has("gamma2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("gamma2 in cache when it should not be")
	}

	inCache, err = testBadgerCache.Has("who")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("who not in cache when it should be")
	}

}
