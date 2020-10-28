package main

import (
	"container/list"
	"fmt"
	"time"
)

var bandera bool
var lista list.List

type Proceso struct {
	Id     uint64
	Status bool
	Canal  chan uint
}

func (p *Proceso) start() {
	i := uint(0)
	for {
		if bandera == true && p.Status == true {
			//p.Canal <- i
			fmt.Printf("id %d: %d", p.Id, i)
			fmt.Println("\n")
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}


func main() {
	var op int64
	var i uint64
	for {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Eliminar proceso")
		fmt.Println("0) Salir")
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("---Se agrego el proceso----")
			p := Proceso{Id: i, Status: true, Canal: make(chan uint)}
			go p.start()
			lista.PushBack(p)
			i++
		case 2:
			fmt.Println("---------Procesos----------")
			bandera = true
		case 3:
			fmt.Println("--------Eliminar----------")
		case 4:
			fmt.Println("-----Adios-----")
			return
		default:
			fmt.Println("Error")
		}
	}
}
