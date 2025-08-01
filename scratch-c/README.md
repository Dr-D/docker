# Minimal C hello world

Uses the scratch image with hello world in 'C' programming language.  
Need to make the file static so that we do not need any external libraries.  
Also use strip to further reduce size of object.

```shell
docker build -t hello-world-scratch .
docker run hello-world-scratch
```
