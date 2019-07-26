package main

import "fmt"

// displayMenu displays the main menu
func displayMenu() {
	fmt.Println(".-------------------------------------------.")
	fmt.Println("|          JSONPLACEHOLDER PHOTOS           |")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("|                                           |")
	fmt.Println("| 1. Listar todas las fotos.                |")
	fmt.Println("| 2. Buscar una foto por ID.                |")
	fmt.Println("| 3. Crear una foto.                        |")
	fmt.Println("| 4. Actualizar una foto.                   |")
	fmt.Println("| 5. Eliminar una foto.                     |")
	fmt.Println("|___________________________________________|")
}

// setUpData receives the entries for a Photo struct
func setUpPhotoStruct(id uint) (p Photo) {
	p.ID = id

	fmt.Printf("Ingrese el ID del álbum -> ")
	fmt.Scanln(&p.AlbumID)

	fmt.Printf("Ingrese el título de la foto -> ")
	fmt.Scanln(&p.Title)

	fmt.Printf("Ingrese la dirección URL de la foto -> ")
	fmt.Scanln(&p.URL)

	fmt.Printf("Ingrese el dirección URL de la miniatura de la foto -> ")
	fmt.Scanln(&p.ThumbnailURL)
	return
}

// showPhoto shows a photo table
func showPhoto(p Photo) {
	fmt.Println("\n_______________________________________________________________________")
	fmt.Println(".------------------------------------------------------------------.")
	fmt.Printf("|                            PHOTO %v                               \n", p.ID)
	fmt.Println("|------------------------------------------------------------------.")
	fmt.Printf("|       ID           |   \t\t%v   \n", p.ID)
	fmt.Printf("|    ALBUMID         |   \t\t%v   \n", p.AlbumID)
	fmt.Printf("|     TITLE          |   %s   \n", p.Title)
	fmt.Printf("|    URLPHOTO        |   %s   \n", p.URL)
	fmt.Printf("| THUMBNAILPHOTO     |   %s   \n", p.ThumbnailURL)
	fmt.Println("|__________________________________________________________________.")
}

// showPhotos shows all the photo tables
func showPhotos() {
	for _, p := range store {
		showPhoto(p)
	}
}

func findPhoto(idPhoto uint) (int, bool) {
	for k, p := range store {
		if p.ID == idPhoto {
			return k, true
		}
	}
	return -1, false
}
