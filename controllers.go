package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var baseURL = "https://jsonplaceholder.typicode.com/photos"

var store photosStore
var lastID uint

func init() {
	data := SendRequest("GET", baseURL, nil)

	// Fill the Store
	err := json.Unmarshal(data, &store)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last ID
	for _, p := range store {
		if p.ID > lastID {
			lastID = p.ID
		}
	}
}

// GetAllPhotos returns all the photos from the Data source
func GetAllPhotos() {
	infoPhotosQuantity()
	time.Sleep(time.Second * 5)
	showPhotos()
}

// GetPhotoByID returns a photo from the data source
func GetPhotoByID() {
	var idPhoto uint

	fmt.Printf("Ingrese el ID de la foto a buscar -> ")
	fmt.Scanln(&idPhoto)

	if photoKey, wasFound := findPhoto(idPhoto); wasFound {
		fmt.Println("\n| \tFOTO ENCONTRADA CON ÉXITO\t |")
		showPhoto(store[photoKey])
		return
	}
	photoNotFound(idPhoto)
}

// CreatePhoto creates a new photo
func CreatePhoto() {
	lastID++
	photo := setUpPhotoStruct(lastID)
	store = append(store, photo)
	fmt.Println("\n| \tFOTO CREADA CON ÉXITO\t |")
	showPhoto(photo)
}

// UpdatePhoto updates a photo by ID
func UpdatePhoto() {
	var idPhoto uint

	fmt.Printf("Ingrese el ID de la foto a actualizar -> ")
	fmt.Scanln(&idPhoto)

	if photoKey, wasFound := findPhoto(idPhoto); wasFound {
		store[photoKey] = setUpPhotoStruct(idPhoto)
		fmt.Println("\n| \tFOTO ACTUALIZADA\t |")
		showPhoto(store[photoKey])
		return
	}
	photoNotFound(idPhoto)
}

// DeletePhoto deletes a photo by ID
func DeletePhoto() {
	var idPhoto uint
	fmt.Printf("Ingrese el ID de la foto -> ")
	fmt.Scanln(&idPhoto)

	if photoKey, wasFound := findPhoto(idPhoto); wasFound {
		lastID--
		store = append(store[:photoKey], store[photoKey+1:]...)
		fmt.Println("\n| \tFOTO ELIMINADA CON ÉXITO\t |")
		return
	}
	photoNotFound(idPhoto)
}
