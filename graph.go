package asciigraph

type Graph interface {
	AddNode(nodeID string)
	AddEdge(nodeOneID, nodeTwoID string)
}

type TestGraph struct {
	Nodes map[string]bool
	Edges map[TestEdge]bool
}

type TestEdge struct {
	nodeOne, nodeTwo string
}

func NewTestGraph() *TestGraph {
	return &TestGraph{
		Nodes: make(map[string]bool),
		Edges: make(map[TestEdge]bool),
	}
}

func (g *TestGraph) AddNode(nodeID string) {
	g.Nodes[nodeID] = true
}

func (g *TestGraph) AddEdge(nodeOneID, nodeTwoID string) {
	e1 := TestEdge{nodeOneID, nodeTwoID}
	g.Edges[e1] = true

	e2 := TestEdge{nodeTwoID, nodeOneID}
	g.Edges[e2] = true
}
