package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}
type SessionManager struct {
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: map[string]Session{},
		mutex:    sync.RWMutex{},
	}
}
func (sm *SessionManager) StartSession(userID string) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sessionID := uuid.New().String()

	newSession := Session{
		UserID:    userID,
		ID:        sessionID,
		ExpiresAt: time.Now().Add(time.Minute * 2),
	}
	sm.sessions[newSession.ID] = newSession

	return sessionID
}
func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	session, ok := sm.sessions[sessionID]
	if !ok || time.Now().After(session.ExpiresAt) {
		return Session{}, false
	}
	return session, true
}
