package handlers

import (
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/iho/gumroad/auth"
	"github.com/nfnt/resize"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Uploading File")
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		log.Println(err)
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}

	defer file.Close()
	log.Print(fmt.Printf("Uploaded File: %+v\n", handler.Filename))
	log.Print(fmt.Printf("File Size: %+v\n", handler.Size))
	log.Print(fmt.Printf("MIME Header: %+v\n", handler.Header))

	saveName := path.Join("/tmp", "ihor", path.Base(handler.Filename))
	savef, err := os.Create(saveName)
	if err != nil {
		// Failed to create file on server, handle err
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer savef.Close()

	if _, err := io.Copy(savef, file); err != nil {
		log.Println("Error")
	}
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	user, ok := auth.ForContext(r.Context())
	response := []byte("fuck you")
	fmt.Println(user)
	if !ok {
		w.Write(response)
		return
	}

	log.Println(w, "Uploading File")
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		log.Println(err)
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}

	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err)
	}
	file.Close()

	log.Print("Uploaded File: %+v\n", handler.Filename)
	log.Print("File Size: %+v\n", handler.Size)
	log.Print("MIME Header: %+v\n", handler.Header)
	m := resize.Resize(600, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Print(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	saveName := path.Join("./", "images", path.Base(handler.Filename))
	savef, err := os.Create(saveName)
	if err != nil {
		// Failed to create file on server, handle err
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer savef.Close()

	if _, err := io.Copy(savef, file); err != nil {
		log.Print("Error")
	}
}
