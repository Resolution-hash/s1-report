FROM golang:1.22-bullseye

# Устанавливаем временную зону
ENV TZ=Europe/Moscow
RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo "Europe/Moscow" > /etc/timezone \
    && apt-get clean && rm -rf /var/lib/apt/lists/*


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps

RUN go build -o main ./app/main.go

CMD ["./main"]
