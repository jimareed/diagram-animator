# diagram-animator
Under development

### build & run
go build
./diagram-animator

### build & run container
```
docker build --tag image2diagram-image .
docker run --name image2diagram -p 8080:8080 -d image2diagram-image
```

### clean up docker
```
docker stop image2diagram
docker rm image2diagram
docker rmi image2diagram-image
```

### run tests
```
go test ./diagram
```