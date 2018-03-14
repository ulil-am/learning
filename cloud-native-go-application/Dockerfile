FROM golang:1.7.5-alpine

MAINTAINER Ulil Amri Kamili

ENV SOURCES /home/ulil/go/src/learning/cloud-native-go-application/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "cloud-native-go-application" ]