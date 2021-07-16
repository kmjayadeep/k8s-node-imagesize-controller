package main

import (
  "os"

  "github.com/urfave/cli/v2"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/rest"
  log "github.com/sirupsen/logrus"
)

func main() {
  app := &cli.App{
    Name: "Controller",
    Usage: "A simple kubernetes controller",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name: "config",
        Usage: "kubeconfig path for usage outside k8s cluster",
      },
    },
  }

  app.Action = func(c *cli.Context) error {
    config := c.String("config")
    log.WithField("config", config).Info("starting controller")

    _, err := getClient(config)

    if err != nil {
      return err
    }

    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
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
