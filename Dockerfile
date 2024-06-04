FROM golang:1.22.3 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/stresser -o stresser

FROM centos
WORKDIR /app
COPY --from=build /app/cmd/stresser/stresser ./
ENTRYPOINT ["./stresser"]