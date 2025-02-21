package async

import (
	"testing"
	"time"
)

func TestAsyncCache_SetAndGet(t *testing.T) {
	cache := NewAsyncCache[string](time.Second)

	// Set a value
	cache.Set("username", "john_doe", 5*time.Second)

	// Get the value
	value, found := cache.GetWithCheck("username")
	if !found {
		t.Errorf("Expected to find key 'username', but it was not found")
	}
	if value != "john_doe" {
		t.Errorf("Expected value 'john_doe', but got '%s'", value)
	}
}

func TestAsyncCache_Expiration(t *testing.T) {
	cache := NewAsyncCache[string](10 * time.Millisecond)

	// Set a value with short TTL
	cache.Set("temp", "expired_value", 20*time.Millisecond)

	// Ensure value is present before expiry
	time.Sleep(10 * time.Millisecond)
	_, found := cache.GetWithCheck("temp")
	if !found {
		t.Errorf("Expected 'temp' key to be present before expiry")
	}

	// Wait for expiration
	time.Sleep(20 * time.Millisecond) // Ensure cleanup goroutine has run
	_, found = cache.GetWithCheck("temp")
	if found {
		t.Errorf("Expected 'temp' key to be expired, but it was found")
	}
}

func TestAsyncCache_GetWithoutCheck(t *testing.T) {
	cache := NewAsyncCache[string](time.Second)

	// Set a value
	cache.Set("name", "Alice", 0) // No expiration

	// Get value without checking expiration
	value, found := cache.GetWithoutCheck("name")
	if !found {
		t.Errorf("Expected key 'name' to be found")
	}
	if value != "Alice" {
		t.Errorf("Expected value 'Alice', but got '%s'", value)
	}
}

func TestAsyncCache_Delete(t *testing.T) {
	cache := NewAsyncCache[string](time.Second)

	// Set and delete a key
	cache.Set("to_delete", "value", 0)
	cache.Delete("to_delete")

	// Ensure key is deleted
	_, found := cache.GetWithCheck("to_delete")
	if found {
		t.Errorf("Expected key 'to_delete' to be deleted, but it was found")
	}
}

func TestAsyncCache_Clear(t *testing.T) {
	cache := NewAsyncCache[string](time.Second)

	// Add multiple keys
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	cache.Clear()

	// Ensure cache is empty
	_, found1 := cache.GetWithCheck("key1")
	_, found2 := cache.GetWithCheck("key2")
	if found1 || found2 {
		t.Errorf("Expected cache to be empty after Clear()")
	}
}
