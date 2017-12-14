package backend

import (
	"sync"

	"github.com/bluemir/zumo/datatype"
	"github.com/sirupsen/logrus"
)

// UserAgentManager handle user agent
type UserAgentManager struct {
	backend *backend

	agents map[string]*UserAgentList
	lock   sync.RWMutex
}

// NewUserAgentManager is
func NewUserAgentManager(b *backend) *UserAgentManager {
	return &UserAgentManager{
		backend: b,
		agents:  map[string]*UserAgentList{},
	}
}

// Register is
func (m *UserAgentManager) Register(name string, ua UserAgent) {
	m.lock.Lock()
	defer m.lock.Unlock()

	logrus.Debugf("[UserAgentManager:Register] name :%s", name)

	if _, ok := m.agents[name]; !ok {
		m.agents[name] = &UserAgentList{}
	}
	m.agents[name].AppendAgent(ua)
}

// Unregister is
func (m *UserAgentManager) Unregister(name string, ua UserAgent) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.agents[name]; !ok {
		// TODO return error or Warn
		return
	}
	// cleanup empty user agent list
	if m.agents[name].isEmpty() {
		delete(m.agents, name)
	}
	m.agents[name].RemoveAgent(ua)
}

func (m *UserAgentManager) get(name string) *UserAgentList {
	m.lock.RLock()
	defer m.lock.RUnlock()

	// TODO check nil
	return m.agents[name]
}

// DeliverJoin is
func (m *UserAgentManager) DeliverJoin(name string, channelID string) {
	m.get(name).DeliverJoin(channelID)
}

// DeliverLeave is
func (m *UserAgentManager) DeliverLeave(name string, channelID string) {
	m.get(name).DeliverLeave(channelID)
}

// DeliverMessage is
func (m *UserAgentManager) DeliverMessage(name string, channelID string, message datatype.Message) {
	m.get(name).DeliverMessage(channelID, message)
}

type UserAgentList struct {
	agents []UserAgent
	lock   sync.RWMutex
}

func (list *UserAgentList) AppendAgent(a UserAgent) {

	list.lock.Lock()
	defer list.lock.Unlock()

	list.agents = append(list.agents, a)

}
func (list *UserAgentList) RemoveAgent(a UserAgent) {

	list.lock.Lock()
	defer list.lock.Unlock()

	for i, agent := range list.agents {
		if agent == a {
			list.agents = append(list.agents[:i], list.agents[i+1:]...)
		}
	}
}
func (list *UserAgentList) DeliverJoin(channelID string) {
	if list == nil {
		return
	}
	list.lock.RLock()
	defer list.lock.RUnlock()

	for _, agent := range list.agents {
		agent.OnJoinChannel(channelID)
	}
}
func (list *UserAgentList) DeliverLeave(channelID string) {
	if list == nil {
		return
	}
	list.lock.RLock()
	defer list.lock.RUnlock()

	for _, agent := range list.agents {
		agent.OnLeaveChannel(channelID)
	}
}
func (list *UserAgentList) DeliverMessage(channelID string, message datatype.Message) {
	if list == nil {
		return
	}
	list.lock.RLock()
	defer list.lock.RUnlock()

	for _, agent := range list.agents {
		agent.OnMessage(channelID, message)
	}
}
func (list *UserAgentList) isEmpty() bool {
	return len(list.agents) == 0
}
