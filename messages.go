package main

import (
	"fmt"
)

// infoPhotosQuantity show the number of photos stored
func infoPhotosQuantity() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Printf("!     SE ENCONTRARON %v FOTOS\n", len(store))
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

// photoNotFound display a photo not found message
func photoNotFound(id uint) {
	fmt.Println(".-------------------------------------.")
	fmt.Printf("| NO SE ENCONTRÓ LA FOTO CON EL ID %v |\n", id)
	fmt.Println(".-------------------------------------.")
}
