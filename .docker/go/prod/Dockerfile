FROM golang:1.13 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG build_path
RUN cd $build_path && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app/main .

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main ./
RUN chmod +x ./main
ENTRYPOINT ./main
EXPOSE 8000