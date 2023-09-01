package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handlerFun)
	fmt.Println("Starting the server on 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func handlerFun(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("GET params were:", r.URL.Query())

	requestVars := r.URL.Query()

	imgUri := requestVars.Get("img")
	requestWidth := requestVars.Get("w")
	requestHeight := requestVars.Get("h")
	data := make(map[string]string)

	if imgUri != "" && requestWidth != "" && requestHeight != "" {

		width, _ := strconv.ParseInt(requestWidth, 10, 64)
		height, _ := strconv.ParseInt(requestHeight, 10, 64)

		src, mimeType := getRemoteImageSrc(imgUri)

		nrgba := imaging.CropAnchor(src, int(width), int(height), imaging.Center)

		imageBase64, _ := imageByteToBase64(nrgba)

		var base64ImageString string
		switch mimeType {
		case "image/jpeg":
			base64ImageString += "data:image/jpeg;base64,"
		case "image/png":
			base64ImageString += "data:image/png;base64,"
		}
		base64ImageString += imageBase64

		data["base64_image"] = base64ImageString
		data["message"] = "Status Created"

		data["status"] = strconv.Itoa(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func getRemoteImageSrc(imgUri string) (image.Image, string) {
	_bytes, _ := getRemoteFile(imgUri)
	mimeType := http.DetectContentType(_bytes)

	reader := bytes.NewReader(_bytes)

	switch mimeType {
	case "image/png":
		_image, _ := png.Decode(reader)
		return _image, mimeType
	case "image/jpeg":
		_image, _ := jpeg.Decode(reader)
		return _image, mimeType
	}

	return nil, ""

}

func imageByteToBase64(dst *image.NRGBA) (string, error) {
	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	if err := imaging.Encode(foo, dst, imaging.JPEG); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func getRemoteFile(uri string) ([]byte, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("Cookie", "name=pat")
	resp, _ := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return _bytes, nil
}
