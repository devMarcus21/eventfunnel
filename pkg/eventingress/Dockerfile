From golang:1.18-alpine

COPY . /go/src/github.com/devMarcus21/eventfunnel/
WORKDIR /go/src/github.com/devMarcus21/eventfunnel/pkg/eventingress/

RUN go build -o /docker-ingress

CMD [ "/docker-ingress" ]

EXPOSE 80