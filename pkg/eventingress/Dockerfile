From golang:1.18-alpine

WORKDIR /ingress

COPY go.mod ./

COPY *.go ./

RUN go build -o /docker-ingress

CMD [ "/docker-ingress" ]

EXPOSE 80