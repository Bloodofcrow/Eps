package main

import (
	"fmt"
)

//Estructura del paciente.
type Node struct {
	Value int
	paciente Paciente
}

//Definición de la estructura: Paciente.
type Paciente struct{
	Nombre string
	Cedula int
	Hora int
	Diagnostico string
	Eps string
}

//Espacio para la definicipon de variables.
var(
	Famisanar = 0
	Compensar = 0
)

func (n *Node) String() string {
	return fmt.Sprint(n.Value, " -> " ,n.paciente.Nombre, " fue atendido y su diagnostico es ", n.paciente.Diagnostico, "\n")
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	if (s.nodes[s.count].paciente.Eps == "Famisanar"){
			Famisanar = Famisanar + 1
	}else if (s.nodes[s.count].paciente.Eps == "Compensar"){
			Compensar = Compensar + 1
	}
	return s.nodes[s.count]
}


func main() {
  fmt.Println("Pila")
	s := NewStack()
	s.Push(&Node{1,Paciente{"Juan", 12345678, 1800, "Reservado.", "Famisanar"}})
	s.Push(&Node{2,Paciente{"Pedro", 12345678, 1800, "Reservado.", "Famisanar"}})
	s.Push(&Node{3,Paciente{"Miguel", 12345678, 1800, "Reservado.", "Compensar"}})
	fmt.Println(s.Pop(), s.Pop(), s.Pop())
	fmt.Println("Se atendieron ", Famisanar, " pacientes de Famisanar.")
	fmt.Println("Se atendieron ", Compensar, " pacientes de Compensar.")
}
