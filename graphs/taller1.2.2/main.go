package main

import (
	"flag"
	"gjae/graph-mesh/algos"
	"gjae/graph-mesh/graph"
)

func main() {

	cols := flag.Int("col", 0, "Indicar la cantidad de columnas para la malla.")
	rows := flag.Int("row", 0, "Indica la cantidad de filas que tendra la malla")
	flag.Parse()

	g := graph.NewGraph(*cols, *rows)
	g.BuildMesh()
	algos.PrimMST(g)
}
