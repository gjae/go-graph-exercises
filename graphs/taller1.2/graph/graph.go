package graph

import (
	"fmt"
	"math/rand"
	"time"
)

type Edge struct {
	v      int
	w      int
	weight float32
}

type GraphMesh struct {
	AdjList map[int][]*Edge
	size    int
}

/**
* Crea un objeto grafo, usando una implementación de lista
* de adyacencia usando un map con una clave entera y de valor una lista
* de aristas
 */
func NewGraph(size int) *GraphMesh {

	vTotal := size * size
	list := make(map[int][]*Edge)
	for i := 0; i < vTotal; i++ {
		list[i] = []*Edge{}
	}

	return &GraphMesh{size: size, AdjList: list}
}

/*
* Agrega una arista entre un origen y un destino
* Se presupone que es un grafo no dirigido por lo que
* se agrega la arista en ambos sentidos
* */
func (g *GraphMesh) AddEdge(origin, dest int, weight float32) {
	g.AdjList[origin] = append(g.AdjList[origin], &Edge{v: origin, w: dest, weight: weight})
	g.AdjList[dest] = append(g.AdjList[dest], &Edge{v: dest, w: origin, weight: weight})
}

func (g *GraphMesh) AdjacentsOf(vertex int) []*Edge {
	if _, ok := g.AdjList[vertex]; ok {
		return g.AdjList[vertex]
	}
	return []*Edge{}
}

/*
* Se crea la malla con las formulas:
* - la formula (i * g.size) + j da el vertice actual
* - la formula (i * g.size) + j + 1 da el vertice siguiente al vertice actual
* - la formula (i * g.size) + j + g.size aplica las aristas verticales de la malla
* Ejemplo con un grafo 3 * 3:
* 0 - 1 - 2
* |	  |   |
* 3 - 4 - 5
* |   |   |
* 6 - 7 - 8
* */
func (g *GraphMesh) BuildMesh() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			currentVertex := (i * g.size) + j
			if j < g.size-1 {
				g.AddEdge(currentVertex, currentVertex+1, rand.Float32())
			}
			if i < g.size-1 {
				g.AddEdge(currentVertex, currentVertex+g.size, rand.Float32())
			}
		}
	}
}

func (g *GraphMesh) Print() {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			currentVertex := (i * g.size) + j
			adVertex := g.AdjList[currentVertex]
			numAdj := len(adVertex)
			if numAdj == 2 {
				if i == 0 && j == 0 {
					fmt.Print("┌")
				} else if i == 0 && j == g.size-1 {
					fmt.Print("┐\n")
				} else if i == g.size-1 && j == 0 {
					fmt.Print("└")
				} else if i == g.size-1 && j == g.size-1 {
					fmt.Print("┘\n")
				}
			} else if numAdj == 3 {
				if j == 0 {
					fmt.Print("├")
				} else if j == g.size-1 {
					fmt.Print("┤\n")
				} else if i == 0 {
					fmt.Print("┬")
				} else {
					fmt.Print("┴")
				}
			} else {
				fmt.Print("┼")
			}
		}
	}
}

func (e *Edge) PrintEdge() {
	fmt.Printf("Desde %d hasta %d. Peso %f\n", e.v, e.w, e.weight)
}
