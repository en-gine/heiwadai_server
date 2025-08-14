package user

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Simple in-memory implementation for testing
type testMemoryRepository struct {
	data map[string][]byte
	mu   sync.Mutex
}

func newTestMemoryRepository() *testMemoryRepository {
	return &testMemoryRepository{
		data: make(map[string][]byte),
	}
}

func (m *testMemoryRepository) Get(key string) *[]byte {
	m.mu.Lock()
	defer m.mu.Unlock()
	if val, ok := m.data[key]; ok {
		return &val
	}
	return nil
}

func (m *testMemoryRepository) Set(key string, value []byte, expire time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
	
	// Simulate expiration
	go func() {
		time.Sleep(expire)
		m.mu.Lock()
		delete(m.data, key)
		m.mu.Unlock()
	}()
}

func (m *testMemoryRepository) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *testMemoryRepository) SetNX(key string, value []byte, expire time.Duration) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if _, exists := m.data[key]; exists {
		return false
	}
	
	m.data[key] = value
	
	// Simulate expiration
	go func() {
		time.Sleep(expire)
		m.mu.Lock()
		delete(m.data, key)
		m.mu.Unlock()
	}()
	
	return true
}

func TestCheckinConcurrentRequests(t *testing.T) {
	memRepo := newTestMemoryRepository()
	
	// Test concurrent SetNX calls
	const numGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	successCount := 0
	var mu sync.Mutex
	
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			
			key := fmt.Sprintf("checkin:%s:%s", "user123", "qr456")
			if memRepo.SetNX(key, []byte("processing"), 5*time.Second) {
				mu.Lock()
				successCount++
				mu.Unlock()
				
				// Simulate processing time
				time.Sleep(100 * time.Millisecond)
				
				// Clean up
				memRepo.Delete(key)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Only one goroutine should succeed
	assert.Equal(t, 1, successCount, "Only one concurrent request should succeed")
}

func TestCheckinKeyExpiration(t *testing.T) {
	memRepo := newTestMemoryRepository()
	
	key := "checkin:user123:qr456"
	
	// First attempt should succeed
	assert.True(t, memRepo.SetNX(key, []byte("processing"), 100*time.Millisecond))
	
	// Second attempt should fail immediately
	assert.False(t, memRepo.SetNX(key, []byte("processing"), 100*time.Millisecond))
	
	// Wait for expiration
	time.Sleep(150 * time.Millisecond)
	
	// After expiration, should succeed again
	assert.True(t, memRepo.SetNX(key, []byte("processing"), 100*time.Millisecond))
}

func TestCheckinWithDelete(t *testing.T) {
	memRepo := newTestMemoryRepository()
	
	key := "checkin:user123:qr456"
	
	// First attempt should succeed
	assert.True(t, memRepo.SetNX(key, []byte("processing"), 5*time.Second))
	
	// Delete the key (simulating completed request)
	memRepo.Delete(key)
	
	// Next attempt should succeed immediately
	assert.True(t, memRepo.SetNX(key, []byte("processing"), 5*time.Second))
}