package main

import (
	"fmt"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

const localEndpoint = "unix:///var/run/docker.sock"

var client *docker.Client

// Connect to Docker
func init() {
	c, err := docker.NewClient(localEndpoint)
	if err != nil {
		panic(err)
	}
	client = c
}

func getAllImages() []docker.APIImages {
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	return imgs
}

func head() {
	fmt.Println("Docker")
	fmt.Println("---")

}

func main() {
	head()

	fmt.Println("Images")
	fmt.Println("---")
	for _, img := range getAllImages() {
		fmt.Println(strings.Split(img.RepoTags[0], ":")[0])
	}
}
