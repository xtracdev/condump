containerbin:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o condump .
	docker build -t xtracdev/condump:latest .

push:
	docker push xtracdev/condump:latest
