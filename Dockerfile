FROM golang:1.20.2 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o simple-server

FROM build AS test
RUN make test

FROM scratch AS final

# Copy binary created in "build" stage
COPY --from=build /app/simple-server .
EXPOSE 8080

ENTRYPOINT ["/simple-server"]
