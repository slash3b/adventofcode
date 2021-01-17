package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
	We are going to build a graph using adjacency list
	Our graph is undirected
	// todo probably I have to add "visited" property
*/

type adjacencyList map[string]*Vertex

type Graph struct {
	vertices adjacencyList
}

func NewGraph() *Graph {
	verMap := make(adjacencyList)

	return &Graph{
		vertices: verMap,
	}
}

func (g *Graph) AddVertex(k string) *Vertex {
	if _, exists := g.vertices[k]; !exists {
		vertex := &Vertex{key: k}

		vertex.adjacent = make(map[string]int)
		g.vertices[k] = vertex
	}

	return g.vertices[k]
}

func (g *Graph) GetVertex(k string) *Vertex {
	if vertex, exists := g.vertices[k]; exists {
		return vertex
	}

	return g.AddVertex(k)
}

func (g *Graph) AddEdge(from, to string, distance int) {
	aVertex := g.GetVertex(from)
	bVertex := g.GetVertex(to)

	aVertex.Tie(to, distance)
	bVertex.Tie(from, distance)
}

func (g *Graph) FindShortestPath(v *Vertex, distance int, visited map[string]int) int {
	fmt.Println("start")
	fmt.Println("going to ", v.key)

	for vertexName, distValue := range v.adjacent {
		if _, exists := visited[vertexName]; exists {
			continue
		}

		// mark as visited
		visited[vertexName]++

		distance += distValue

		fmt.Println("going to ", vertexName)
		return g.FindShortestPath(g.GetVertex(vertexName), distance, visited)
	}

	// now all were visited so we can reset visited
	fmt.Println("end")
	return distance
}

type Vertex struct {
	key      string // name
	adjacent map[string]int
}

func (v *Vertex) Tie(vertexName string, distance int) {
	if _, exists := v.adjacent[vertexName]; !exists {
		v.adjacent[vertexName] = distance
	}
}

func main() {
	graph := prepare("/home/slash3b/Projects/aoc/2015/go/day9/input")

	for _, vertex := range graph.vertices {
		val := graph.FindShortestPath(vertex, 0, make(map[string]int))
		fmt.Println("Distance: ", val)
	}

}

func prepare(s string) *Graph {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	graph := NewGraph()

	for _, v := range inputs {
		chunks := strings.Split(v, "to")
		aNode := strings.TrimSpace(chunks[0])

		chunks = strings.Split(chunks[1], "=")
		bNode := strings.TrimSpace(chunks[0])

		distance, err := strconv.Atoi(strings.TrimSpace(chunks[1]))
		if err != nil {
			panic(err)
		}

		graph.AddEdge(aNode, bNode, distance)
	}

	return graph
}
