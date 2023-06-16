From golang:1.18-alpine

COPY . /go/src/github.com/devMarcus21/eventfunnel/
WORKDIR /go/src/github.com/devMarcus21/eventfunnel/pkg/subscriptionfanout/

RUN go build -o /docker-subscription-fanout

CMD [ "/docker-subscription-fanout" ]

EXPOSE 80