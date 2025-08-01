FROM gcc:latest AS build
COPY hello.c /hello.c
RUN gcc -Os -static hello.c -o hello
RUN strip /hello && \
    ls -la /hello

FROM scratch
COPY --from=build /hello /hello
ENTRYPOINT ["/hello"]
