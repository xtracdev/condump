# Condump

Periodically dump containers and their health to standard out. Can be
useful for keeping an eye on what's going on in say an ECS cluster

## Running it

<pre>
docker run --volume=/var/run:/var/run:rw xtracdev/condump:latest
</pre>

## Dependencies

go get github.com/samalba/dockerclient
go get github.com/Sirupsen/logrus
