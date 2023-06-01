FROM alpine:3.14

RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing \
  bash git datamash jq curl make jshon perl build-base ca-certificates htmlq

# setting setzer configs
ENV SETZER_CACHE_EXPIRY=-1 \
  SETZER_MIN_MEDIAN=4

# Installing setzer
COPY ./libexec/setzer/ /root/setzer

ENV PATH="/root/setzer:${PATH}"

# Copy sources
COPY . /app

WORKDIR /app

CMD ["/bin/bash"]
