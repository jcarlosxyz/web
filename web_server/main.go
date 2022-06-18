//devemos corerlos co go run .  <---se compilan y se ejecutan todos los archivos de goland del directorio
//Pero antes debemos correr la siguinte sentencia
//go mod init anyname

package main

func main() {

	server := NewServer(":3000")
	server.Listen()

}
