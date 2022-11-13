package main

import (
	"fmt"
	"math/rand"
	"time"
)

func latidos_de_perro() {
	var alleatorio int
	var limite int
	limite = 100
	for i := 0; i <= limite; i = i + alleatorio {
		alleatorio = rand.Intn(5-1) + 1
		time.Sleep(500 * time.Millisecond)
		fmt.Println("El perro tiene ", alleatorio, " latidos por segundo")
	}
}

func latidos_de_gato() {
	var alleatorio int
	var limite int
	limite = 100
	for i := 0; i <= limite; i = i + alleatorio {
		alleatorio = rand.Intn(5-1) + 1
		time.Sleep(500 * time.Millisecond)
		fmt.Println("El gato tiene ", alleatorio, " latidos por segundo")
	}
}

func main() {
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	go latidos_de_perro()
	go latidos_de_gato()
	var wait string
	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Scanln(&wait)

}
