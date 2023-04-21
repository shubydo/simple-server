FROM golang:1.20.3 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o simple-server

FROM build AS test
RUN make test

FROM scratch AS final
ARG PORT=8080

# Copy binary created in "build" stage
COPY --from=build /app/simple-server /simple-server
EXPOSE $PORT

ENTRYPOINT ["/simple-server", "start", "--port", "$PORT"]