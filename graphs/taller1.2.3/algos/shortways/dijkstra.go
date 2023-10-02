package shortways

import (
	"container/heap"
	"gjae/graph-mesh/graph"
)

type PriorityQueue []Queue

type Queue struct {
	Source int
	Dist   float32
}

type DijkstraAlg struct {
	Mesh  graph.GraphMesh
	EdgeA []*graph.Edge
	DistA []float32
	Queue PriorityQueue
}

const INF = 999999999999999999999.0

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Las aristas con menor peso tienen mayor prioridad
	return pq[i].Source < pq[j].Source
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Queue))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Has(w int) (bool, int) {
	for i := 0; i < pq.Len(); i++ {
		if (*pq)[i].Source == w {
			return true, i
		}
	}

	return false, -1
}

func (g *DijkstraAlg) Relax(v Queue) {
	source := v.Source

	for _, edge := range g.Mesh.AdjacentsOf(source) {
		w := edge.Oposed(source)

		if g.DistA[w] > g.DistA[source]+edge.Weight() {
			g.DistA[w] = g.DistA[source] + edge.Weight()
			g.EdgeA[w] = edge
			if ok, k := g.Queue.Has(w); ok {
				g.Queue[k].Dist = g.DistA[w]
			} else {
				heap.Push(&g.Queue, Queue{Source: w, Dist: g.DistA[w]})
			}
		}
	}
}

func (g *DijkstraAlg) DistanceOf(v int) float32 {
	return g.DistA[v]
}

func (g *DijkstraAlg) ExistRoute(source int) bool {
	return g.DistA[source] < INF
}

func NewDijkstra(mesh *graph.GraphMesh) *DijkstraAlg {
	dist := make([]float32, mesh.Size()*mesh.Size())

	for i := 0; i < mesh.Size()*mesh.Size(); i++ {
		dist[i] = INF
	}

	return &DijkstraAlg{
		Mesh:  *mesh,
		EdgeA: make([]*graph.Edge, mesh.Size()*mesh.Size()),
		DistA: dist,
		Queue: make(PriorityQueue, 0),
	}
}

func RunDijkstra(graph *Graph, origin int) {
	d := NewDijkstra(graph.Mesh)

	heap.Init(&d.Queue)
	d.EdgeA[origin] = nil
	d.DistA[origin] = 0
	heap.Push(&d.Queue, Queue{Source: origin, Dist: 0})

	for d.Queue.Len() > 0 {
		v := heap.Pop(&d.Queue)
		d.Relax(v.(Queue))
		graph.RelaxCounter(DIJKSTRA)
	}

}
