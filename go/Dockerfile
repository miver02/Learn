FROM golang:1.22

WORKDIR /go/src/app

RUN apt-get update && apt-get install -y \
	vim \
	git \
	&& rm -rf /var/lib/apt/lists/*

COPY . /go/src/app

CMD ["bash"]
