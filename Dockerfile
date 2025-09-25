FROM golang:1.22
RUN apt-get update && apt-get install -y \
    build-essential \
    gcc \
    vim \
    git \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /root/learn
ENV CGO_ENABLED=1

# Python deps
COPY requirements.txt /tmp/requirements.txt
RUN pip3 install --no-cache-dir -r /tmp/requirements.txt || true
CMD ["bash"]