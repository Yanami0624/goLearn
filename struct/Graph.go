package data_struct

import (
	. "Gocodes/algorithm"
	"fmt"
	"math/rand"
)

type Graph[T Comp] struct {
	nodes     []*Node[T]
	edges     []map[int]int
	Is_single bool
}

func BuildGraph[T Comp](n int, complexity float64, is_sigle bool) *Graph[T] {
	g := new(Graph[T])
	g.Is_single = is_sigle
	for range n {
		g.AddNode()
	}
	judge := func() bool {
		return rand.Float64() < complexity
	}
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			if judge() {
				g.AddEdge(i, j)
			}
		}
	}
	return g
}

func BuildWeightedGraph[T Comp](n int, complexity float64, is_sigle bool, maxWeight int) *Graph[T] {
	g := new(Graph[T])
	g.Is_single = is_sigle
	for range n {
		g.AddNode()
	}
	judge := func() bool {
		return rand.Float64() < complexity
	}
	weight := func() int {
		if maxWeight <= 1 {
			return 1
		}
		return rand.Intn(maxWeight) + 1
	}
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			if judge() {
				g.AddWeightedEdge(i, j, weight())
			}
		}
	}
	return g
}

func (g *Graph[T]) Print() {
	bound := func() {
		for range g.Size() {
			fmt.Print("--")
		}
		fmt.Println()
	}
	bound()
	for n, edge := range g.edges {
		fmt.Printf("%d: ", n)
		for k, v := range edge {
			fmt.Printf("%d(%d) ", k, v)
		}
		fmt.Println()
	}
	bound()
}

func (g *Graph[T]) AddNode() {
	g.nodes = append(g.nodes, new(Node[T]))
	g.edges = append(g.edges, make(map[int]int))
}

func (g *Graph[T]) AddEdge(src, dst int) {
	g.AddWeightedEdge(src, dst, 1)
}

func (g *Graph[T]) AddWeightedEdge(src, dst, weight int) {
	g.edges[src][dst] = weight
	if !g.Is_single {
		g.edges[dst][src] = weight
	}
	src %= g.Size()
	dst %= g.Size()
	g.nodes[src].AddWeightedEdge(g.nodes[dst], weight)
	if !g.Is_single {
		g.nodes[dst].AddWeightedEdge(g.nodes[src], weight)
	}
}

func (g *Graph[T]) Size() int {
	return len(g.nodes)
}

func (g *Graph[T]) TopoSort() (ret []int) {
	if !g.Is_single {
		return
	}
	n := g.Size()
	degree := make([]int, n)
	for _, es := range g.edges {
		for k := range es {
			degree[k]++
		}
	}

	q := Queue[int]{}
	for src := range degree {
		if degree[src] == 0 {
			q.Push(src)
		}
	}
	for q.Size() != 0 {
		cur := q.Front()
		for dst := range g.edges[cur] {
			degree[dst]--
			if degree[dst] == 0 {
				q.Push(dst)
			}
		}
		q.Pop()
		ret = append(ret, cur)
	}
	return
}

func (g *Graph[T]) ExistRing() bool {
	if !g.Is_single {
		return true
	}
	n := g.Size()
	degree := make([]int, n)
	for _, es := range g.edges {
		for k := range es {
			degree[k]++
		}
	}

	q := Queue[int]{}
	for src := range degree {
		if degree[src] == 0 {
			q.Push(src)
		}
	}
	ret := make([]int, 0)
	for q.Size() != 0 {
		cur := q.Front()
		for dst := range g.edges[cur] {
			degree[dst]--
			if degree[dst] == 0 {
				q.Push(dst)
			}
		}
		q.Pop()
		ret = append(ret, cur)
	}
	return len(ret) != n
}

func (g *Graph[T]) Dijkstra() int {
	n := g.Size()
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0

	q := Heap[dijkstraItem]{
		Greater: func(a, b dijkstraItem) bool {
			return a.dis < b.dis
		},
	}
	q.Push(dijkstraItem{node: 0})
	for q.Len() != 0 {
		cur := q.Pop()
		if cur.dis != dis[cur.node] {
			continue
		}
		if cur.node == n-1 {
			return cur.dis
		}
		for dst, weight := range g.edges[cur.node] {
			nextDis := cur.dis + weight
			if dis[dst] != -1 && dis[dst] <= nextDis {
				continue
			}
			dis[dst] = nextDis
			q.Push(dijkstraItem{node: dst, dis: nextDis})
		}
	}
	return -1
}

type dijkstraItem struct {
	node int
	dis  int
}

type Node[T Comp] struct {
	val      T
	neibours map[*Node[T]]int
}

func (n *Node[T]) AddEdge(dst *Node[T]) {
	n.AddWeightedEdge(dst, 1)
}

func (n *Node[T]) AddWeightedEdge(dst *Node[T], weight int) {
	if n.neibours == nil {
		n.neibours = make(map[*Node[T]]int)
	}
	n.neibours[dst] = weight
}
