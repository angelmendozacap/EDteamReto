package main

import (
	"fmt"
	"time"
)

func main() {

	var opt uint8
	for {
		fmt.Println("_______________________________________________________________________________")
		displayMenu()
		fmt.Printf("\t*\n\t*Elija una opción -> ")
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
			fmt.Println("| -_- ESA OPCIÓN NO ESTÁ DISPONIBLE -_- |")
			break
		}
		time.Sleep(time.Second * 4)
	}
}
