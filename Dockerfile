FROM golang:latest
WORKDIR /go
ADD . /go
ENV PORT=:8080
#CMD ["go", "run", "main.go"]
CMD ["go", "build", "-o", "e-search"]
CMD ["./e-search"]
