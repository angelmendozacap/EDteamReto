package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var baseURL = "https://jsonplaceholder.typicode.com/photos"

// GetAllPhotos returns all the photos from the Data source
func GetAllPhotos() {
	photos := []Photo{}
	data := SendRequest("GET", baseURL, nil)

	err := json.Unmarshal(data, &photos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(photos)
}

// GetPhotoByID returns a photo from the data source
func GetPhotoByID() {
	photo := Photo{}
	var idPhoto uint

	fmt.Printf("Ingrese el ID de la foto -> ")
	fmt.Scanln(&idPhoto)

	// Endpoint
	URL := fmt.Sprintf("%s/%v", baseURL, idPhoto)
	// Getting the data
	data := SendRequest("GET", URL, nil)

	err := json.Unmarshal(data, &photo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(photo)
}

// CreatePhoto creates a new photo
func CreatePhoto() {
	photo := Photo{}

	setUpPhotoStruct(&photo)

	data := SendRequest("POST", baseURL, photo)
	fmt.Println(string(data))
}

// UpdatePhoto updates a photo by ID
func UpdatePhoto() {
	photo := Photo{}
	var idPhoto uint

	fmt.Printf("Ingrese el ID de la foto -> ")
	fmt.Scanln(&idPhoto)

	URL := fmt.Sprintf("%s/%v", baseURL, idPhoto)

	setUpPhotoStruct(&photo)

	data := SendRequest("PUT", URL, photo)
	fmt.Println(string(data))
}

// DeletePhoto deletes a photo by ID
func DeletePhoto() {
	var idPhoto uint
	fmt.Printf("Ingrese el ID de la foto -> ")
	fmt.Scanln(&idPhoto)

	URL := fmt.Sprintf("%s/%v", baseURL, idPhoto)

	data := SendRequest("DELETE", URL, nil)
	fmt.Println(string(data))
}
