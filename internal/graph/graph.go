package graph

import (
	"errors"
	"fmt"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node struct {
	name  string
	pos   rl.Vector2
	edges map[*Node]struct{}
}

func GetNodePos(n *Node) rl.Vector2 {
	return n.pos
}

func CreateNode(name string, pos rl.Vector2) *Node {
	return &Node{
		name:  name,
		pos:   pos,
		edges: make(map[*Node]struct{}),
	}
}

func GetNeighbors(node *Node, graph *Graph, gridWidth, gridHeight int) {
	x, y := node.pos.X, node.pos.Y
	for n := range graph.nodes {
		if n.pos.X == x+1 && n.pos.Y == y {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x-1 && n.pos.Y == y {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x && n.pos.Y == y-1 {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x && n.pos.Y == y+1 {
			addEdgeToNode(graph, node, n)
		}
		if n.pos.X == x-1 && n.pos.Y == y-1 {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x+1 && n.pos.Y == y-1 {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x-1 && n.pos.Y == y+1 {
			addEdgeToNode(graph, node, n)
		}

		if n.pos.X == x+1 && n.pos.Y == y+1 {
			addEdgeToNode(graph, node, n)
		}
	}
}

type Graph struct {
	nodes map[*Node]struct{}
}

func GetNodes(g *Graph) []*Node {
	ret := make([]*Node, 0)

	for n := range g.nodes {
		ret = append(ret, n)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].name < ret[j].name
	})

	return ret
}

func AddNode(g *Graph, n *Node) error {
	if g.nodes == nil {
		g.nodes = make(map[*Node]struct{})
	}

	if _, ok := g.nodes[n]; ok {
		return errors.New("Node is already in the graph")
	}

	g.nodes[n] = struct{}{}

	return nil
}

func addEdgeToNode(g *Graph, n *Node, ns ...*Node) error {
	if _, ok := g.nodes[n]; !ok {
		return errors.New("Node to add edges to is not in the graph")
	}

	for _, nToAdd := range ns {
		if _, ok := g.nodes[nToAdd]; !ok {
			return errors.New("Node being added as edge is not in the graph")
		}

		if n.edges == nil {
			n.edges = make(map[*Node]struct{})
		}

		if _, ok := n.edges[nToAdd]; ok {
			return errors.New("This edge is already added to this node")
		}

		n.edges[nToAdd] = struct{}{}
		if nToAdd.edges == nil {
			nToAdd.edges = make(map[*Node]struct{})
		}
		nToAdd.edges[n] = struct{}{}
	}

	return nil
}

func PrintGraph(g *Graph) {
	for n := range g.nodes {
		edgeNames := make([]string, 0)

		for e := range n.edges {
			edgeNames = append(edgeNames, e.name)
		}

		fmt.Println(n.name, edgeNames)
	}
}
