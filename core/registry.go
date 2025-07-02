package core

import "sync"

type NodeRegistry struct {
	mu    sync.Mutex
	nodes map[string]*Node
}

func NewRegistry() *NodeRegistry {
	return &NodeRegistry{nodes: make(map[string]*Node)}
}

func (r *NodeRegistry) Register(node *Node) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nodes[node.ID] = node
}

func (r *NodeRegistry) List() []*Node {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := []*Node{}
	for _, n := range r.nodes {
		result = append(result, n)
	}
	return result
}
