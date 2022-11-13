package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Servidor struct {
	direccion string
	descarga  string
}

func (s Servidor) Descarga() {
	var alleatorio int
	var limite int
	limite = 100 - 5
	for i := 0; i <= 100; i = i + alleatorio {
		alleatorio = rand.Intn(5-1) + 1
		fmt.Println("El servidor ", s.direccion, " esta pasando paquetes de ", s.descarga, " al ", i, "%")
		time.Sleep(500 * time.Millisecond)
		if i >= limite {
			fmt.Println(s.descarga, " Se ha descargado al 100%")
			break
		}
	}
}

func main() {
	descarga1 := Servidor{
		direccion: "192.152.218.12",
		descarga:  "La era de hielo",
	}

	descarga2 := Servidor{
		direccion: "198.126.548.246",
		descarga:  "Crystal Glasses - not in love",
	}

	t1 := time.Now()
	t2 := t1.Add(time.Second)
	go descarga1.Descarga()
	go descarga2.Descarga()
	var wait string
	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Scanln(&wait)

}
