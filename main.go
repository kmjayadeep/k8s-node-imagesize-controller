package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var imageSizes map[string]uint64

func main() {
	app := &cli.App{
		Name:  "Controller",
		Usage: "A simple kubernetes controller",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "kubeconfig path for usage outside k8s cluster",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		config := c.String("config")
		log.WithField("config", config).Info("starting controller")

		client, err := getClient(config)

		if err != nil {
			return err
		}

		err = watchNodes(client)

		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func watchNodes(clientset *kubernetes.Clientset) error {
	imageSizes = make(map[string]uint64)
	api := clientset.CoreV1().Nodes()
	watcher, err := api.Watch(context.Background(), metav1.ListOptions{})

	if err != nil {
		return err
	}

	ch := watcher.ResultChan()

	for event := range ch {
		node, ok := event.Object.(*coreV1.Node)
		if !ok {
			watcher.Stop()
			return fmt.Errorf("couldn't cast to nodelist")
		}

		checkImageSize(node)
	}

	return nil
}

func checkImageSize(node *coreV1.Node) {

	var size uint64 = 0

	for _, img := range node.Status.Images {
		size = size + uint64(img.SizeBytes)
	}

	if imageSizes[node.Name] != size {
		log.Infof("Image size changed for node %s. Old: [%v], New: [%v]", node.Name, humanize.Bytes(imageSizes[node.Name]), humanize.Bytes(size))
		imageSizes[node.Name] = size
	}

}

func getClient(config string) (*kubernetes.Clientset, error) {
	var cfg *rest.Config
	var err error
	if config == "" {
		// in cluster config
		cfg, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	cfg, err = clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(cfg)
}
