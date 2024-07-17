
FROM ubuntu:24.04

LABEL maintainer="nawaphon"

RUN useradd -ms /bin/bash nawaphon

WORKDIR /home/nawaphon

COPY ./build/gin-play /bin/

# Command to run the executable
CMD ["gin-play"]