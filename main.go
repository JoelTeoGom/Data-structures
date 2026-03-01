package main

import (
	"data-structures/mapa"
	"fmt"
)

func main() {
	// s := stack.Build[int](0, 5)

	// s.Push(10)
	// s.Push(20)
	// s.Push(30)

	// fmt.Println(s.Elements)

	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Elements)
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())

	///queue

	// q := queue.Build[int](0, 5)

	// q.Queue(10)
	// q.Queue(20)
	// q.Queue(30)

	// fmt.Println(q.Elements)

	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Elements)
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())

	m := mapa.Build(8)

	m.Set("nombre", "Joel")
	m.Set("edad", 25)
	m.Set("ciudad", "Barcelona")
	m.Set("lenguaje", "Go")

	m.Print()

	fmt.Println("---")

	v, ok := m.Get("nombre")
	fmt.Println("Get nombre:", v, ok)

	v, ok = m.Get("noExiste")
	fmt.Println("Get noExiste:", v, ok)

	v, ok = m.Get("edad")
	fmt.Println("Get edad:", v, ok)

	v, ok = m.Get("ciudad")
	fmt.Println("Get ciudad:", v, ok)

	m.Delete("ciudad")
	fmt.Println("--- despues de Delete ciudad ---")
	m.Print()

	v, ok = m.Get("ciudad")
	fmt.Println("Get ciudad:", v, ok)

	v, ok = m.Get("lenguaje")
	fmt.Println("Get lenguaje:", v, ok)

}
