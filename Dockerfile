FROM golang:1.11.1

WORKDIR /go/src/hot_reload_docker
COPY . .
ENV GO111MODULE=on
ENV LINE_CHANNEL_SECRET=ad320b00ff56657a9a7b7088bad2d7eb
    LINE_CHANNEL_TOKEN='oXkSkZpPjPFLMD5D9WegcidPrwWghQ4A3BvTUzg4wk0eYcUTumPAYiRUK708LLmgrY+4paJGbfwksI4GpCA3k/g/RAtXGRayqaBes9P47yZVC9psxcpKZvfVIMgKRnuy7loFeneehEmo8rpopw6AfgdB04t89/1O/w1cDnyilFU='
    PORT=9000

EXPOSE 8083
RUN go get github.com/pilu/fresh

CMD ["fresh"]
