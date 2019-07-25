package main

import (
	"fmt"
)

func main() {

	var opt uint8

	displayMenu()

	for {
		fmt.Printf("\n\tElija una opción -> ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			GetAllPhotos()
		case 2:
			GetPhotoByID()
		case 3:
			CreatePhoto()
		case 4:
			UpdatePhoto()
		case 5:
			DeletePhoto()
		default:
			fmt.Println("Esa opción no esta disponible")
			break
		}
	}
}
