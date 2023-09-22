package main

import (
	"flag"
	"gjae/graph-mesh/algos"
	"gjae/graph-mesh/graph"
)

func main() {

	size := flag.Int("size", 0, "Indicar un tamaño para la malla. Sera una malla N*N")
	flag.Parse()

	g := graph.NewGraph(*size)
	g.BuildMesh()

	algos.PrimMST(g)
}
