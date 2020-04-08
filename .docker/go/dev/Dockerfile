FROM golang:1.13

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG install_compile_deamon
ARG install_goose

RUN if [ $install_goose = "on" ] ; then go get -u github.com/pressly/goose/cmd/goose ; echo "Installing goose is turned off" ; fi
RUN if [ $install_compile_deamon = "on" ] ; then go get github.com/githubnemo/CompileDaemon ; echo "Installing CompileDeamon is turned off" ; fi

COPY . .
