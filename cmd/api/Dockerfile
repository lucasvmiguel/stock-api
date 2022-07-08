##
## BUILD
## 
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN ls

RUN go build /app/cmd/api/main.go

##
## RUN
##
FROM alpine

WORKDIR /

COPY --from=build /app/main /main
COPY --from=build /app/.env /.env

RUN ls

ENTRYPOINT ["/main"]