package main

import (
	"gjae/graph-mesh/graph"
)

func main() {
	g := graph.NewGraph(4)
	g.BuildMesh()
	g.Print()
}
