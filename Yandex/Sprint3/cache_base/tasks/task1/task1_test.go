package main

import (
	"testing"
	"time"
)

func TestSessionManager_StartSession(t *testing.T) {
	sm := NewSessionManager()
	userID := "user123"

	sessionID := sm.StartSession(userID)
	if sessionID == "" {
		t.Error("Expected session ID to be non-empty")
	}

	session, ok := sm.GetSession(sessionID)
	if !ok {
		t.Error("Expected session to exist")
	}

	if session.UserID != userID {
		t.Errorf("Expected session user ID to be %s, got %s", userID, session.UserID)
	}
}

func TestSessionManager_GetSession(t *testing.T) {
	sm := NewSessionManager()
	userID := "user123"

	sessionID := sm.StartSession(userID)

	session, ok := sm.GetSession(sessionID)
	if !ok {
		t.Error("Expected session to exist")
	}

	if session.UserID != userID {
		t.Errorf("Expected session user ID to be %s, got %s", userID, session.UserID)
	}

	// Test session expiry
	time.Sleep(time.Minute * 3) // Wait for session to expire
	_, ok = sm.GetSession(sessionID)
	if ok {
		t.Error("Expected session to be expired")
	}
}
