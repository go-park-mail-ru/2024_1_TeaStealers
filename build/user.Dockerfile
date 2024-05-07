FROM golang:1.21.0-alpine AS builder

COPY . /github.com/go-park-mail-ru/2024_1_TeaStealers/
WORKDIR /github.com/go-park-mail-ru/2024_1_TeaStealers/

RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/auth/main.go


FROM scratch AS runner

WORKDIR /docker-cian-user/

COPY --from=builder /github.com/go-park-mail-ru/2024_1_TeaStealers/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_1_TeaStealers/config config/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip
EXPOSE 80 443

ENTRYPOINT ["./.bin"]