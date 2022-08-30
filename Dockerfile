FROM golang:1.18-alpine AS build

WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /app/main ./cmd/web 

FROM alpine:3.16.2
WORKDIR /
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD [ "/app/main" ]