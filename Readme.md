# Tech Challenge Winter Semester 18-19

## Overview
This is the technical documentation for the tech challenge project *Building Sphere* by the Triplex group. This repository hosts every subsystem of the project. The codebase is written in [Golang](https://golang.org/). It uses docker as the development infrastructure. To work on this project, you need to have both Golang and Docker installed on your system.

> To see the demo, you can visit http://141.40.254.162 , although that server will be undeployed in the future.

> For screenshots, you can visit the images folder.

## Project Structure
- **images:** Stores the screenshosts of the system.
- **reinvent:** This folder has the codebase for central advertisement system. It creates a server that allows third party comes to add advertisement to system.
- **Sphere:** This folder hosts the main dashboard served to the users. It has 3 functions. It collects real time data from the microcontrollers in the sockets, it processes ads and maps it to certain users and it runs the main dashboard server.
- **Socket:** Socket is the code to run on microcontrollers in the power sockets. Every 1 minute, it sends the time and the socket ID to the server. If the socket is not being used currently, it doesn't consume electricity.
- **Test:** Hosts the script for random data generation.
- **db.sql:** Our MySQL database structure. It creates a db called `test`, creates tables used by the system and inserts some mock data. You can directly import this to your database.
- **docker-compose.yaml:** It builds the individual projects and runs them together. 
- **makefile:** It compiles individual projects. Just to compile projects run `make`, to run them `make docker`. You need to have *docker* and *golang* installed on your system.

# Used technologies
## Communication
This project uses **MQTT** for the communication between sockets and server. Sockets publish to the topic `my-topic` and server is subscribed to that topic. An MQTT broker is distributed with the *docker-compose.yaml* in this project. Docker handles networking itself, thus the hostname for the MQTT server is given as `mqtt`, where the mqtt broker is mapped to.
## Database
The system uses `MySQL` database for data storage. There are 6 tables.

- **ads:** Stores ads data. Every ad has a type and usage information which is used to compare ads with power usage data.
- **measurements:** Stores measurements sent by the sockets. Every socket has a unique id and the time that the measurement was collected. 
- **notifs:** Stores notifications. Notifications are simple texts that concern every inhabitant of the building.
- **promotions:** Stores promotions created by ReInvent.
- **repairs:** Stores repair requests from house owners to reinvent. Requests have a status field that shows their completion.
- **sockets:** Stores microcontrolller data used by sockets. Every microcontroller has a unique id, a type and a status. The type is e.g. [Entertainment, Lighting, Cooking, Charging]

To reach the database you can go to port `8080` on your deployment. An additional image called `adminer` is deployed by the system to allow GUI operations on the database.

## Server
Sphere server listens on port `4000`. It has a javascript chart library for rendering charts. It uses both ajax and server-side rendering for data retrieval.

Reinvent server listens on port `5000`.

# Hardware
- The reinvent code will be running on the cloud. In the demo, it runs on the LRZ Cloud.

- The sphere will be running locally in the building. It is currently deployed to a Raspberry Pi, but more processing will require a better hardware.

- The socket code will be running on ESP wifi modules. Golang can be compiled to almost any environment, thus transporting only the binary is okey.

![An Example ESP module](https://asset.conrad.com/media10/isa/160267/c1/-/de/1656367_BB_00_FB/entwickler-platine-sbc-nodemcu-esp32.jpg?x=520&y=520= "An example ESP Module")

# Running
- Make sure you have Golang and docker installed.
- Run `make docker`. If Golang cannot find dependencies, go to sphere folder and run `go get`.
- Once every service starts, go to `localhost:8080` and connect to db. Username is `root`, password is `example`. Import `db.sql` to the database. 
- Go to `localhost:4000` to see user dashboard.
- Go to `localhost:5000` to see reinvent dashboard.

> In docker compose , socket deployment is disabled not to pollute the system with data everytime a new deployment is made. You populate data, you can either enable it in docker-compose, or run `go run main.go` in test folder.
