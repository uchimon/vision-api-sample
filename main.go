// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

type Output struct {
	Image     string   `json:"image"`
	Labels    []Label  `json:"labels"`
	Landmarks []string `json:"landmarks"`
}
type Label struct {
	Description string  `json:"description"`
	Score       float32 `json:"score"`
	Topicality  float32 `json:"topicality"`
}

func main() {
	ctx := context.Background()
	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	images := readFile("./imglist.txt")
	outpus := make([]Output, 0, len(images))
	for _, uri := range images {
		fmt.Println(uri)
		outpus = append(outpus, visionAPIDataByURI(ctx, client, uri))
	}

	b, err := json.Marshal(outpus)
	if err != nil {
		log.Fatal("json error", err)
	}
	write(b, "./out.json")
	fmt.Println(string(b))
}

func visionAPIDataByURI(ctx context.Context, client *vision.ImageAnnotatorClient, uri string) Output {

	//##  Sets the name of the image file to annotate.
	//filename := "./test.jpg"

	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatalf("Failed to read file: %v", err)
	// }
	// defer file.Close()
	// image, err := vision.NewImageFromReader(file)
	// if err != nil {
	// 	log.Fatalf("Failed to create image: %v", err)
	// }

	// ## from url
	image := vision.NewImageFromURI(uri)

	labels, err := client.DetectLabels(ctx, image, nil, 30)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	outLabels := make([]Label, 0, len(labels))

	fmt.Println("Labels:")
	for _, label := range labels {
		fmt.Println(label.Description)
		fmt.Println(label)
		l := Label{
			Description: label.Description,
			Score:       label.Score,
			Topicality:  label.Topicality,
		}
		outLabels = append(outLabels, l)
	}

	landmarks, err := client.DetectLandmarks(ctx, image, nil, 30)
	if err != nil {
		log.Fatalf("Failed to detect lnads: %v", err)
	}

	fmt.Println("Landmarks:")
	outLands := make([]string, 0, len(landmarks))
	for _, landmark := range landmarks {
		fmt.Println(landmark.Description)
		fmt.Println(landmark)
		outLands = append(outLands, landmark.Description)
	}
	return Output{
		Image:     uri,
		Labels:    outLabels,
		Landmarks: outLands,
	}
}

func readFile(path string) (lines []string) {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines
}

func write(val []byte, path string) {
	os.Remove(path)
	err := ioutil.WriteFile(path, val, 0666)
	if err != nil {
		log.Fatal("error ", err)
	}
}
