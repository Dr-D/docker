# Go Static Executable in scratch

Creates a minimal image(approx 5.64MB when built) that checks for a file and then returns a 200 response on localhost:9090.
Useful for checking if a container is ready/up in kubernetes cluster.

We have to use an alternative to standard c lib as this is not available in the scratch image.
This is why musl is used to allow static linkage.

```shell
cd scratch-go
docker build -t go-static-executable .
docker run  --name=static-go -p9090:9090 -v ./:/opt/ go-static-executable
```

Open new shell:
```shell
curl localhost:9090 # should return false
false 
```
We have mapped local directory on host to /opt on the container.
```shell
touch ready.txt
curl localhost:9090
true # should return true
```
