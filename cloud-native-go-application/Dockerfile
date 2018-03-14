FROM golang:1.7.5-alpine

MAINTAINER Ulil Amri Kamili

ENV GOPATH /home/ulil/go
ENV SOURCES /home/ulil/go/src/learning/cloud-native-go-application/
WORKDIR /home/ulil/go/src/learning/cloud-native-go-application

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV PORT 8080
EXPOSE 8080

RUN chmod +x cloud-native-go-application

ENTRYPOINT ["./cloud-native-go-application"]