package main

import "fmt"

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
func setUpPhotoStruct(p *Photo) {
	fmt.Printf("Ingrese el ID del álbum -> ")
	fmt.Scanln(p.AlbumID)

	fmt.Printf("Ingrese el título de la foto -> ")
	fmt.Scanln(p.Title)

	fmt.Printf("Ingrese la dirección URL de la foto -> ")
	fmt.Scanln(p.URL)

	fmt.Printf("Ingrese el dirección URL de la miniatura de la foto -> ")
	fmt.Scanln(p.ThumbnailURL)
}
