package main //Nombre del packete
//package cualquier nombre

import "fmt" //Importar librerias necesarias para ejecutar el programa

func main() {
	var mensaje string = "Hola mundo"
	mensajeFacil := "Hola mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)
	//Comentarios!!!
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)
	var c int = 10
	d := 3
	fmt.Println(c / d)
	x := true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)
	privada()
	Publica()
	imprimirHola()
}

func privada() {
	fmt.Println("Ejecutar logica que no necesita ser exportada")
}

func Publica() {
	fmt.Println("Logica que auiero exportar")
}

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
