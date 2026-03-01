package mapa

import "fmt"

type Map struct {
	Element []Bucket
	nElems  int
}
type Bucket struct {
	Nodes []Node
}
type Node struct {
	Key   string
	Value any
}

func Build(size int) *Map {
	if size <= 0 {
		size = 16
	}
	elem := make([]Bucket, size)
	return &Map{Element: elem}
}

func (m *Map) Size() int {
	return m.nElems
}

func hashKey(key string, bucketSize int) int {
	h := 0
	for _, ch := range key {
		h = h*31 + int(ch)
	}
	if h < 0 {
		h = -h
	}
	return h % bucketSize
}

func resize(m *Map) {
	if m.nElems >= len(m.Element) {
		newSize := len(m.Element) * 2
		newElement := make([]Bucket, newSize)
		for _, bucket := range m.Element {
			for _, node := range bucket.Nodes {
				index := hashKey(node.Key, newSize)
				newElement[index].Nodes = append(newElement[index].Nodes, node)
			}
		}
		m.Element = newElement
	}
}

func (m *Map) Set(key string, value any) bool {
	if len(m.Element) > 0 {
		index := hashKey(key, len(m.Element))
		m.Element[index].Nodes = append(m.Element[index].Nodes, Node{
			Key:   key,
			Value: value,
		})
		m.nElems++
		resize(m)
		return true
	}
	return false
}

func (m *Map) Get(key string) (any, bool) {
	if len(m.Element) > 0 {
		index := hashKey(key, len(m.Element))
		for _, node := range m.Element[index].Nodes {
			if node.Key == key {
				return node.Value, true
			}
		}
	}
	return nil, false
}

func (m *Map) Print() {
	fmt.Printf("Map (size: %d, buckets: %d)\n", m.nElems, len(m.Element))
	for i, bucket := range m.Element {
		if len(bucket.Nodes) > 0 {
			fmt.Printf("  Bucket[%d]: ", i)
			for _, node := range bucket.Nodes {
				fmt.Printf("    %q => %v -", node.Key, node.Value)
			}
		}
		fmt.Println()
	}
}

func (m *Map) Delete(key string) bool {
	if len(m.Element) > 0 {
		index := hashKey(key, len(m.Element))
		for i, node := range m.Element[index].Nodes {
			if node.Key == key {
				m.Element[index].Nodes = append(m.Element[index].Nodes[:i], m.Element[index].Nodes[i+1:]...)
				m.nElems--
				return true
			}
		}
	}
	return false
}
