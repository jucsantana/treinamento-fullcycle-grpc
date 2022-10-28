FROM golang

VOLUME [ "/data", "/usr/src/app" ]

WORKDIR /usr/src/app

RUN apt update -y &&\
    apt install protobuf-compiler -y &&\
    go mod init github.com/jucsantana/fc2-grpc &&\
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go &&\
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc &&\
    go get google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install github.com/ktr0731/evans@latest

CMD ["tail","-f","/dev/null"]



