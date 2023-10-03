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

func (g *Graph) Print() {
	g.Mesh.Print()
	var entry, exit int

	for k, _ := range g.Mesh.AdjList {
		if len(g.Mesh.AdjacentsOf(k)) == 1 {
			if entry == 0 {
				entry = k
			} else if exit == 0 {
				exit = k
			}
		}

		if entry != 0 && exit != 0 {
			break
		}
	}
	RunDijkstra(g, entry, exit)

	fmt.Printf("\nDijkstra: %d\n", g.dijkstra)
}
