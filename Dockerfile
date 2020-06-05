FROM ubuntu:20.10

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y
RUN apt-get install -yq texlive-full
RUN apt-get install -y git

RUN apt update -y
RUN apt install zsh -y
RUN apt install curl -y
RUN curl -JLO https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh && chmod +x install.sh && ./install.sh

WORKDIR /docs