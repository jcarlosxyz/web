//Aplicacion de concurrencia
//
//
//

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	fmt.Println(inicio, "Inicio de la aplicacion")
	canal := make(chan string)
	servidores := []string{
		"https://platzi.com/",
		"https://www.google.com.mx/",
		"https://www.youtube.com/",
		"https://twitter.com/",
	}

	i := 0
	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {

			go resvisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	//for i := 1; i < len(servidores); i++ {

	//		fmt.Println(<-canal)
	//}

	fmt.Println("Timpo transcorrido de la aplicacion ", time.Since(inicio))

}

func resvisarServidor(servidor string, canal chan string) {

	_, err := http.Get(servidor)

	if err != nil {

		//fmt.Println("El servidor", servidor, "esta fuera de  servicio")
		canal <- servidor + "No se encuentra activo"
	} else {

		//fmt.Println("El servidor ", servidor, "esta funcionado normalmente")
		canal <- servidor + "Se encuentra activo"
	}

}
