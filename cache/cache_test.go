package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	err := testRedisCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo in cache when it should not be")
	}

	err = testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("expected foo in cache but its not there")
	}
}


func TestRedisCache_Get(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("expected to find value bar in cache but not there")
	}
}

func TestRedisCache_Forget(t *testing.T) {
	err := testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Forget("alpha")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("value in cache when it should not be")
	}
}

func TestRedisCache_Empty(t *testing.T) {
	err := testRedisCache.Set("delta", "gamma")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("delta")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("value in cache when it should be empty")
	}
}

func TestRedisCache_EmptyByMatch(t *testing.T) {
	err := testRedisCache.Set("delta", "gamma")
	if err != nil {
		t.Error(err)
	}
	err = testRedisCache.Set("delta2", "omega")
	if err != nil {
		t.Error(err)
	}
	err = testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.EmptyByMatch("delta")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("delta")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("key in cache when it should not be")
	} 

	inCache, err = testRedisCache.Has("delta2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("key delta2  in cache when it should not be")
	}

	inCache, err = testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}
	if !inCache {
		t.Error("key alpha not in cache when it should be")
	}

}

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"

	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}

	_, err = decode(string(bytes))
	if err != nil {
		t.Error(err)
	}
}
