FROM golang:1.16.2-alpine as builder

COPY . /app

WORKDIR /app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o pechanger ./backend/cmd/pechanger
RUN rm -rf .git &&  find . -type f -not -name pechanger -not -name pechanger_config.json -delete

CMD [ "ls", "-la"   ]


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app . 

CMD [ "./pechanger" ]
