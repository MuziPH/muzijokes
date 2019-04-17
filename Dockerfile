FROM golang:alpine
MAINTAINER Muzi

ENV SOURCES /go/src/jokesapi/

COPY . ${SOURCES}

RUN go get github.com/gorilla/mux
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT jokesapi