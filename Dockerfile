FROM golang
COPY main.go .
RUN go build -o hello_world main.go
EXPOSE 8080
ENTRYPOINT /go/hello_world

#FROM alpine
#COPY --from=hello_world /go/hello_world /
#EXPOSE 8080
#CMD /hello_world
