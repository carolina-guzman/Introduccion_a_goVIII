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
			p.Canal <- i + 1
			go p.print()
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) cleanChanel() { //un canal necesita recibir y dar informacion
	for {
		y := <-p.Canal
		y = y + 1
	}
}

func (p *Proceso) print() {
	for {
		a := <-p.Canal
		fmt.Printf("id %d: %d", p.Id, a)
		fmt.Println()
	}
}

func stop(id uint64) {
	for i := lista.Front(); i != nil; i = i.Next() {
		if i.Value.(*Proceso).Id == id {
			i.Value.(*Proceso).Status = false
			break
		}
	}
}

func main() {
	var op int64
	var i uint64
	bandera = false
	for {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Eliminar proceso")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("---Se agrego el proceso----")
			p := Proceso{Id: i, Status: true, Canal: make(chan uint)}
			go p.start()
			go p.cleanChanel()
			lista.PushBack(p)
			i++
			break
		case 2:
			fmt.Println("---------Procesos----------")
			bandera = true
			var input string
			fmt.Scanln(&input)
			bandera = false
			break
		case 3:
			fmt.Println("--------Eliminar----------")
			var idp uint64
			fmt.Println("ID del proceso: ")
			fmt.Scanln(&idp)
			stop(idp)
			break
		case 0:
			fmt.Println("-----Adios-----")
			return
			break
		default:
			fmt.Println("Error")
			break
		}
	}
}
