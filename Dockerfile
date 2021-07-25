FROM golang:1.16
LABEL maintainer="Deepak Singhvi <deepak.singhvi@gmail.com>"
ENV GO111MODULE=on

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build
RUN go install

ENTRYPOINT ["catalogwf"]
EXPOSE 7442
CMD ["catalogservice"]