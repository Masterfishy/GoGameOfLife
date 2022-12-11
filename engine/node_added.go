package engine

var NodeAdded nodeAdded

type NodeClass string

const (
	Render NodeClass = "Render"
	Living NodeClass = "Living"
)

type NodeAddedPayload struct {
	Class NodeClass
	RenderNode *RenderNode
	LivingNode *LivingNode
}

type nodeAdded struct {
	handlers []interface{ NodeAddedHandler(NodeAddedPayload) }
}

// AddListener adds an event handler for this event
func (c *nodeAdded) AddListener(handler interface{ NodeAddedHandler(NodeAddedPayload) }) {
	c.handlers = append(c.handlers, handler)
}

// Invoke sends out an event with the payload
func (c nodeAdded) Invoke(payload NodeAddedPayload) {
	for _, handler := range c.handlers {
		go handler.NodeAddedHandler(payload)
	}
}
