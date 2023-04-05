FROM golang:1.20 as build

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

ENV PATH="/root/go/bin:${PATH}"
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make proto

RUN make docker-build-server

EXPOSE 8080

CMD ["./bin/kango-server"]

# Final stage
FROM scratch
COPY --from=build /app/bin/kango-server /kango-server
EXPOSE 8080
ENTRYPOINT ["/kango-server"]
