build:
	env GOOS=linux GOARCH=amd64 go build -o socket/socket-linux socket/*.go
	env GOOS=linux GOARCH=amd64 go build -o sphere/sphere-linux sphere/*.go

clean:
	-rm socket/socket-linux
	-rm sphere/sphere-linux
	-rm sphere/sphere
docker:
	make clean
	make build
	sleep 1
	docker-compose build
	docker-compose up