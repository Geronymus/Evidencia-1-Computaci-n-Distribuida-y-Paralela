package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Caballo struct {
	nombre string
	inicio int
	fin    int
}

func (c Caballo) Avanza() {
	var alleatorio int
	var limite int
	limite = c.fin - 5
	for i := c.inicio; i <= c.fin; i = i + alleatorio {
		alleatorio = rand.Intn(5-1) + 1
		time.Sleep(500 * time.Millisecond)
		fmt.Println(c.nombre, " estoy cabalgando en el paso: ", i)
		if i >= limite {
			fmt.Println(c.nombre, " Ha llegado a la meta!")
			break
		}
	}
}

func main() {
	antares := Caballo{
		nombre: "Antares",
		inicio: 0,
		fin:    100,
	}

	aldebaran := Caballo{
		nombre: "Aldebaran",
		inicio: 0,
		fin:    100,
	}

	t1 := time.Now()
	t2 := t1.Add(time.Second)
	antares.Avanza()
	aldebaran.Avanza()
	diff := t2.Sub(t1)
	fmt.Println(diff)
}
