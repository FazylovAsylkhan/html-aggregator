FROM golang:1.23.5-alpine AS builder
WORKDIR /usr/local/src
RUN apk --no-cache add bash git make gcc gettext musl-dev postgresql-client
RUN go env -w GOBIN=/usr/local/bin
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download
COPY app ./
COPY app/sqlc.yaml ./sqlc.yaml
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
RUN go build -o /usr/local/bin/app cmd/main.go

FROM alpine AS runner
RUN apk --no-cache add \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    ca-certificates \
    postgresql-client \ 
    && ln -s /usr/lib/chromium/chrome /usr/bin/google-chrome 
COPY --from=builder /usr/local/bin/app /usr/local/bin/
ENV CHROME_BIN="/usr/bin/google-chrome"
CMD [ "/usr/local/bin/app"  ]