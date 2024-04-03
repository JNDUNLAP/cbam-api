FROM golang:1.22 as base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM base AS development

RUN go install github.com/cosmtrek/air@v1.27.3

WORKDIR /app

CMD ["air"]

FROM base AS production

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM alpine:3.14 AS final

ARG APP_USER

RUN apk --no-cache add ca-certificates && \
    adduser -D "${APP_USER}"

WORKDIR /app

COPY --from=production /app/main .

USER "${APP_USER}"

CMD ["./main"]
