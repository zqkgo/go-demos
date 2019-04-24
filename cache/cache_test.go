package cache

import (
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	c := NewCache(10)
	err := c.Set("", "", 0)
	if err == nil {
		t.Error("should return error but not")
	}
	c.Set("foo", "bar", 1*time.Second)
	v, _ := c.Get("foo")
	if v != "bar" {
		t.Error("should return bar but not")
	}
	<-time.After(1 * time.Second)
	v, _ = c.Get("foo")
	if v != nil {
		t.Error("should return nil but not")
	}
	if c.KeyNum() != 0 {
		t.Error("should return 0 but not")
	}

	c.Set("a", "a", 1*time.Second)
	c.Set("b", "b", 1*time.Second)
	c.Set("c", "c", 1*time.Second)
	c.Set("d", "d", 2*time.Second)
	if c.KeyNum() != 4 {
		t.Error("should return 4 but not")
	}
	<-time.After(1 * time.Second)
	c.Collect()
	if c.KeyNum() != 1 {
		t.Error("should return 1 after Collect() but not")
	}
	v, _ = c.Get("d")
	if v != "d" {
		t.Error("should return d but not")
	}

	c.Del("d")
	c.Set("hello", "world", 10*time.Second)
	c.Set("key", "val", 10*time.Second)
	c.Set("123", "456", 10*time.Second)
	c.Del("hello")
	if c.KeyNum() != 2 {
		t.Error("KeyNum() should return 2 after Del() but not")
	}
	c.Flush()
	if c.KeyNum() != 0 {
		t.Error("KeyNum() should return 0 after flush but not")
	}
}

func TestCollector(t *testing.T) {
	c := NewCache(1)
	if !c.Collector().Running() {
		t.Fatal("collector should be runing but not")
	}
	c.Set("foo", "bar", 1*time.Second)
	c.Set("biz", "buz", 2*time.Second)
	c.Set("hello", "world", 2*time.Second)
	<-time.After(3 * time.Second)
	if c.KeyNum() != 0 {
		t.Fatal("invalid items should all be collected")
	}

	c.Collector().Stop()
	<-time.After(500 * time.Millisecond)
	if c.Collector().Running() {
		t.Fatal("collector should be stopped but not")
	}

	c.Set("foo", "bar", 1*time.Second)
	c.Set("biz", "buz", 2*time.Second)
	c.Set("hello", "world", 2*time.Second)
	<-time.After(3 * time.Second)
	if c.KeyNum() == 0 {
		t.Fatal("invalid items should not be collected because collector is stopped")
	}

	c.Collector().Collect()
	<-time.After(500 * time.Millisecond)
	if c.KeyNum() != 0 {
		t.Fatal("invalid items should all be collected because collector is started")
	}
}

func TestStructVal(t *testing.T) {
	type Dummy struct {
		Name string
	}
	d := Dummy{"Unknown"}
	c := NewCache(1)
	c.Set("k", d, 3 * time.Second)
	v, _ := c.Get("k")
	cv := v.(Dummy)
	if cv.Name != "Unknown" {
		t.Fatal("fail to parse struct type value")
	}
}
