package main

import (
	"time"
	"github.com/samalba/dockerclient"
	log "github.com/Sirupsen/logrus"
)

const (
	MetricsDumpInterval = 5 * time.Second
)



func main() {

	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	c := time.Tick(MetricsDumpInterval)
	for range c {
		//Signal self to dump metrics to stdout
		dumpHealth(docker)
	}
}

func dumpHealth(docker *dockerclient.DockerClient) {
	containers, err := docker.ListContainers(false, false, "")
	if err != nil {
		log.Errorf("Error listing containers: %s", err.Error())
		return
	}

	for _,c := range containers {
		log.Infof("%s %s %s",
			c.Names[0],
			c.Image,
			c.Status,
		)
	}
}