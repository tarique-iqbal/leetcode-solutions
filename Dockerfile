FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive \
    TZ=UTC \
    GO_VERSION=1.24.1 \
    APP_DIR=/home/leetcode

RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /tmp/* /var/tmp/*

RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR ${APP_DIR}

COPY go.mod ./
RUN go mod tidy

COPY . .

CMD ["go", "test", "-v", "./..."]
