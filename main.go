package main

import (
	"fmt"
	"math"
)

type bot interface {
	getGreeting() string
}

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (englishBot) getGreeting() string {
	return "Hello There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
