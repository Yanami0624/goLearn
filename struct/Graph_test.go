package data_struct

import (
	. "Gocodes/algorithm"
	"testing"
)

func TestBuildWeightedGraph(t *testing.T) {
	const (
		n         = 4
		maxWeight = 8
	)
	g := BuildWeightedGraph[MyInt](n, 1, true, maxWeight)

	if got := g.Size(); got != n {
		t.Fatalf("expected graph size %d, got %d", n, got)
	}
	for src, edge := range g.edges {
		if got := len(edge); got != n-1 {
			t.Fatalf("expected node %d to have %d edges, got %d", src, n-1, got)
		}
		if _, ok := edge[src]; ok {
			t.Fatalf("node %d should not have self edge", src)
		}
		for dst, weight := range edge {
			if weight < 1 || weight > maxWeight {
				t.Fatalf("edge %d -> %d weight %d out of range", src, dst, weight)
			}
		}
	}
}

func TestBuildWeightedGraphDefaultWeight(t *testing.T) {
	g := BuildWeightedGraph[MyInt](3, 1, true, 0)

	for src, edge := range g.edges {
		for dst, weight := range edge {
			if weight != 1 {
				t.Fatalf("expected edge %d -> %d default weight 1, got %d", src, dst, weight)
			}
		}
	}
}

func TestGraphDijkstra(t *testing.T) {
	g := new(Graph[MyInt])
	for range 5 {
		g.AddNode()
	}
	g.Is_single = true

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	if got := g.Dijkstra(); got != 3 {
		t.Fatalf("expected shortest distance 3, got %d", got)
	}
}

func TestGraphDijkstraWeighted(t *testing.T) {
	g := new(Graph[MyInt])
	for range 4 {
		g.AddNode()
	}
	g.Is_single = true

	g.AddWeightedEdge(0, 1, 10)
	g.AddWeightedEdge(1, 3, 10)
	g.AddWeightedEdge(0, 2, 1)
	g.AddWeightedEdge(2, 3, 2)

	if got := g.Dijkstra(); got != 3 {
		t.Fatalf("expected weighted shortest distance 3, got %d", got)
	}
}

func TestGraphDijkstraUnreachable(t *testing.T) {
	g := new(Graph[MyInt])
	for range 3 {
		g.AddNode()
	}
	g.Is_single = true
	g.AddEdge(0, 1)

	if got := g.Dijkstra(); got != -1 {
		t.Fatalf("expected unreachable distance -1, got %d", got)
	}
}
