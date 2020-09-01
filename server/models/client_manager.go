package models

import "sync"

type ClientManager struct {
	Clients     map[*Client]bool   // all connection
	ClientsLock sync.RWMutex       // lock
	Users       map[string]*Client // appId+uuid
	UserLock    sync.RWMutex       // lock
	Register    chan *Client
	Login       chan *login
	Unregister  chan *Client
	Broadcast   chan []byte // broadcast
}

// 初始化 (傳入指標)
func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}
