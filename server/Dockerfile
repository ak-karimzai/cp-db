FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app main.go
EXPOSE 8080
CMD ["/app/app"]

# sudo docker build -f Dockerfile . -t cp-db
# sudo docker start --network host -p 8080:8080 cp-db