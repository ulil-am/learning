FROM alpine:3.5

MAINTAINER Ulil Amri Kamili

COPY ./cloud-native-go-application /app/cloud-native-go-application
RUN chmod +x /app/cloud-native-go-application

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/cloud-native-go-application