

build multistage
docker build -t docker-gs-ping:multistage -f Dockerfile.multistage .


run container:
```sh
docker run -p 8080:8080 -it booklibrary
```


# shell
## execute an interactive sh shell on the container.
```sh
docker exec -it mycontainer sh
```

## execute a command on the container.
 ```sh
 docker exec -d mycontainer touch /tmp/execWorks`
 ```



 create an image:
 `docker build -t libraryapp:multistage -f Dockerfile.multistage .`
 run:
 `docker run -p 8080:8080 -it libraryapp:multistage`


 # this project

build:
docker build -t booklibrary .
docker build --build-arg OS=linux --build-arg ARCH=amd64 -t booklibrary_amd64 .

run:
docker run -p 8080:8080 -it booklibrary
rn with environment variables:
docker run -d --env PORT=8082 -p 8080:8082 -it booklibrary
docker run -d --env PORT=8082 --env RUNTIME_SETUP=dev -p 8080:8082 -it booklibrary
pass env file
docker run -d --env-file .env -p 8080:8082 -it booklibrary
docker run -d --env-file .env -p 8080:8082 -it booklibrary_amd64



# docker buildx
create

`docker buildx create --name gobooklibbuilder`
use:

`docker buildx use gobooklibbuilder`

insepect:

`docker buildx inspect gobooklibbuilder`
(this should show Status: inactive)
active:

`docker buildx inspect gobooklibbuilder --bootstrap`

docker build push on repo
`docker buildx build --platform linux/arm64,linux/amd64 -t isrmino/go_book_library . --push`
docker pull
`docker pull isrmino/go_book_library:latest`

docker run -d --env-file .env -p 8080:8082 -it isrmino/booklibrary



# tools used in the course
- delve (debug)
- fluentd
- loggly ( web portal)