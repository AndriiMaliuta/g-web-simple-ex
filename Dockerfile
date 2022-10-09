#syntax=docker/dockerfile:1
FROM golang:1.18-alpine
WORKDIR /app
COPY . .
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
RUN go get .
#COPY *.go ./
RUN go build -o /go-app

# run
#COPY --from=go_build /go-app go-app
#EXPOSE 4001
CMD [ "/go-app" ]