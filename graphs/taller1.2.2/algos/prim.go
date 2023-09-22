package algos

import (
	"container/heap"
	mesh "gjae/graph-mesh/graph"
)

type PriorityQueue []*Edge

type Edge struct {
	edge *mesh.Edge
	Char string
}

type Prim struct {
	Queue   PriorityQueue
	Visited []bool
	TreeRm  []*Edge
}

const INF = 9999999

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Las aristas con menor peso tienen mayor prioridad
	return pq[i].edge.Weight() < pq[j].edge.Weight()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (p *Prim) Visit(g *mesh.GraphMesh, v int) {
	p.Visited[v] = true
	for _, edge := range g.AdjacentsOf(v) {
		if !p.Visited[edge.Oposed()] {
			heap.Push(&p.Queue, &Edge{edge: edge})
		}
	}

}

func PrimMST(graph *mesh.GraphMesh) {
	prim := Prim{
		Visited: make([]bool, graph.Size()*graph.Size()),
		Queue:   make(PriorityQueue, 0),
		TreeRm:  make([]*Edge, graph.Size()*graph.Size()),
	}

	heap.Init(&prim.Queue)
	// Agregar las aristas a la cola con prioridad

	prim.Visit(graph, 0)

	for prim.Queue.Len() > 0 {
		edge := heap.Pop(&prim.Queue)
		v, w := edge.(*Edge).edge.Vertexes()
		if prim.Visited[v] && prim.Visited[w] {
			continue
		}
		prim.TreeRm = append(prim.TreeRm, edge.(*Edge))
		edge.(*Edge).edge.PrintEdge()

		if !prim.Visited[v] {
			prim.Visit(graph, v)
		}
		if !prim.Visited[w] {
			prim.Visit(graph, w)
		}
	}

}
