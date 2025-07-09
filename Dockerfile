FROM golang:alphine

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /github.com/VeryFach/MMD-FILKOM-TINGAL

COPY . .

RUN go build -o github.com/VeryFach/MMD-FILKOM-TINGAL

CMD ["/github.com/VeryFach/MMD-FILKOM-TINGAL/github.com/VeryFach/MMD-FILKOM-TINGAL"]