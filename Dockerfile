FROM golang:1.16-alpine

WORKDIR $GOPATH/src/github.com/DayDzen/k8s-sandbox

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 80

CMD ["k8s-sandbox"]