package event

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"sync"
)

// Event manager
// Handles sending and receiving events for immediate handling
// Generally used for engine functionality, such as user input events, window
// management etc.
type Manager struct {
	listenerMap map[message.Id]map[core.Handle]IEventListenable
	mu          sync.Mutex
	eventQueue  []*QueueItem
	runAsync    bool
}

//Register a new component to listen to an event
func (manager *Manager) Listen(eventName message.Id, component IEventListenable) core.Handle {
	handle := core.NewHandle()
	manager.mu.Lock()
	if _, ok := manager.listenerMap[eventName]; !ok {
		manager.listenerMap[eventName] = make(map[core.Handle]IEventListenable)
	}
	manager.listenerMap[eventName][handle] = component
	manager.mu.Unlock()

	return handle
}

// Runs the event queue in its own go routine
func (manager *Manager) RunConcurrent() {
	// Block double-running
	if manager.runAsync == true {
		return
	}
	manager.runAsync = true
	go func() {
		for manager.runAsync == true {
			manager.mu.Lock()
			queue := manager.eventQueue
			manager.mu.Unlock()

			if len(queue) > 0 {
				// FIFO - ensure dispatch order, and concurrency integrity
				item := queue[0]
				manager.mu.Lock()
				manager.eventQueue = manager.eventQueue[1:]

				// Fire event
				listeners := manager.listenerMap[item.EventName]
				manager.mu.Unlock()
				for _, component := range listeners {
					component.ReceiveMessage(item.Message)
				}
			}
		}
	}()
}

//Remove a listener from listening for an event
func (manager *Manager) Unlisten(eventName message.Id, handle core.Handle) {
	manager.mu.Lock()
	if _, ok := manager.listenerMap[eventName][handle]; ok {
		delete(manager.listenerMap[eventName], handle)
	}
	manager.mu.Unlock()
}

//Fires an event to all listening components
func (manager *Manager) Dispatch(eventName message.Id, message message.IMessage) {
	message.SetType(eventName)
	queueItem := &QueueItem{
		EventName: eventName,
		Message:   message,
	}
	manager.mu.Lock()
	manager.eventQueue = append(manager.eventQueue, queueItem)
	manager.mu.Unlock()
}

// Close the event manager
func (manager *Manager) Unregister() {
	// Ensure async event queue is halted
	manager.runAsync = false
}

var eventManager Manager

func GetEventManager() *Manager {
	if eventManager.listenerMap == nil {
		eventManager.listenerMap = make(map[message.Id]map[core.Handle]IEventListenable)
	}
	return &eventManager
}
