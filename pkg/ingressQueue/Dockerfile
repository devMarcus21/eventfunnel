From golang:1.18-alpine

COPY . /go/src/github.com/devMarcus21/eventfunnel/
WORKDIR /go/src/github.com/devMarcus21/eventfunnel/pkg/ingressQueue/

RUN go build -o /docker-ingress-queue-consumer

EXPOSE 4569

CMD [ "/docker-ingress-queue-consumer" ]