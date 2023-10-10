FROM golang:1.20.3 AS cache

WORKDIR /app
COPY go.* .
RUN go mod download

FROM cache AS build

WORKDIR /app
COPY . .

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -o simple-server

FROM build AS test
RUN make test

FROM scratch AS final

# Copy binary created in "build" stage
COPY --from=build /app/simple-server /simple-server

ENTRYPOINT [ "/simple-server"]
CMD ["start", "-p", "8080"]
