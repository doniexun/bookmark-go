FROM instrumentisto/glide:latest

WORKDIR /app

COPY . /app

RUN glide install
RUN go build -o main .
CMD ["/app/main"]