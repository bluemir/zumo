package backend

import (
	l "sync"

	"github.com/sirupsen/logrus"
)

// UserAgentManager handle user agent
type UserAgentManager struct {
	backend *backend

	agents map[string][]UserAgent
	lock   *l.RWMutex
}

// NewUserAgentManager is
func NewUserAgentManager(b *backend) *UserAgentManager {
	return &UserAgentManager{
		backend: b,
		agents:  map[string][]UserAgent{},
		lock:    &l.RWMutex{},
	}
}

// Register is
func (m *UserAgentManager) Register(name string, ua UserAgent) {
	m.lock.Lock()
	defer m.lock.Unlock()

	logrus.Debugf("[UserAgentManager:Register] name :%s", name)

	m.agents[name] = append(m.agents[name], ua)
}

// Unregister is
func (m *UserAgentManager) Unregister(name string, ua UserAgent) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, a := range m.agents[name] {
		if a == ua {
			m.agents[name] = append(m.agents[name][:i], m.agents[name][i+1:]...)
		}
	}
}

// DeliverJoin is
func (m *UserAgentManager) DeliverJoin(name string, channelID string) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	logrus.Debugf("[UserAgentManager:DeliverJoin] name: %s, channelID: %s", name, channelID)

	for _, a := range m.agents[name] {
		m.backend.channels[channelID].AddListener(a)

		a.OnJoinChannel(channelID)
	}
}

// DeliverLeave is
func (m *UserAgentManager) DeliverLeave(name string, channelID string) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, a := range m.agents[name] {
		// m.backend.channels[channelID].RemoveListener(a)

		a.OnLeaveChannel(channelID)
	}
}
