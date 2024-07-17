
FROM golang:1.22

LABEL maintainer="nawaphon"

RUN useradd -ms /bin/bash nawaphon

WORKDIR /home/nawaphon

RUN go get https://github.com/nawaphonOHM/gin-play

RUN make

EXPOSE 8080

CMD ["gin-play"]