package main

import (
	"fmt"
	"time"

	"github.com/mmdali-dev/easycache/async"
)

// Session represents a user session.
type Session struct {
	UserID   int
	Username string
	Expires  time.Time
}

func main() {
	// Create an async cache with a cleanup interval of 10 seconds.
	sessionCache := async.NewAsyncCache[Session](10 * time.Second)

	// Create a sample session.
	session := Session{
		UserID:   101,
		Username: "john_doe",
		Expires:  time.Now().Add(30 * time.Minute),
	}

	// Store the session with a TTL of 30 minutes.
	sessionCache.Set("session_101", session, 30*time.Minute)

	// Retrieve the session from the cache.
	if cachedSession, found := sessionCache.GetWithCheck("session_101"); found {
		fmt.Printf("✅ Session Found: %+v\n", cachedSession)
	} else {
		fmt.Println("❌ Session not found or expired")
	}

	// Simulate session expiration.
	time.Sleep(2 * time.Second)
	sessionCache.Delete("session_101") // Manually delete the session.

	// Check if the session is still available.
	if _, found := sessionCache.GetWithCheck("session_101"); !found {
		fmt.Println("✅ Session expired and deleted successfully 🧹")
	} else {
		fmt.Println("❌ Unexpected session found")
	}
}
