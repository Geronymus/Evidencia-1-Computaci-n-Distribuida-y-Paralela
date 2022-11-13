package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Caballo struct {
	nombre string
	inicio int
	fin    int
}

// Se pasa este parametro por referencia
func (c Caballo) Avanza(wg *sync.WaitGroup) {
	var alleatorio int
	var limite int
	limite = c.fin - 5
	for i := c.inicio; i <= c.fin; i = i + alleatorio {
		alleatorio = rand.Intn(5-1) + 1
		time.Sleep(500 * time.Millisecond)
		fmt.Println(c.nombre, " estoy cabalgando en el paso: ", i)
		if i >= limite {
			fmt.Println(c.nombre, " Ha llegado a la meta!")
			//Señala que el elemento del grupo de espera ya termino
			defer wg.Done()
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

	var wg sync.WaitGroup
	//Añadir elementos al grupo de espera
	wg.Add(2)
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	//Se les pasa la dirección del grupo de espera
	//es un puntero
	go antares.Avanza(&wg)
	go aldebaran.Avanza(&wg)
	diff := t2.Sub(t1)
	fmt.Println(diff)
	//Bloquea el hilo de ejecución de la goroutine hasta que todos los elementos
	//del WaitGroup esten terminados
	wg.Wait()
}
