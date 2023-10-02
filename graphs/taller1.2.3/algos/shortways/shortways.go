package shortways

import (
	"fmt"
	"gjae/graph-mesh/graph"
)

type Graph struct {
	Mesh     *graph.GraphMesh
	dijkstra int
	bfs      int
	aStar    int
}

const (
	DIJKSTRA = iota
	BFS
	ASTAR
)

func NewShortWayFinder(old *graph.GraphMesh) *Graph {
	return &Graph{
		Mesh: old,
	}
}

func (g *Graph) RelaxCounter(relaxSource int) {
	if relaxSource == DIJKSTRA {
		g.dijkstra++
	} else if relaxSource == BFS {
		fmt.Println("Relajando BFS")
		g.bfs++
	} else {
		fmt.Println("Relajando A*")
		g.aStar++
	}
}

func (g *Graph) Print(source int) {
	g.Mesh.Print()
	RunDijkstra(g, source)

	fmt.Printf("\nDijkstra: %d\n", g.dijkstra)
}
