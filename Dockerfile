FROM arm32v7/golang

RUN mkdir /app

ADD usernameAndPassword.go /app/
ADD first.go /app/

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]