ARG VERSION

FROM golang:1.25-alpine AS build
ARG VERSION

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -ldflags "-X main.version=${VERSION}" -o ./dist/mealplan ./cmd

FROM golang:1.25-alpine AS deploy

COPY --from=build ./app/dist/mealplan ./mealplan

ENTRYPOINT [ "./mealplan" ]