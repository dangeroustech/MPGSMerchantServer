#Clone and build app
FROM golang:1.14.0-buster as GO
WORKDIR /app
ARG TRAVIS_BRANCH=$TRAVIS_BRANCH
ARG BUILD_ENV=$BUILD_ENV
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV PORT=8080
RUN apt update && apt upgrade -y && apt install -y git
RUN git clone https://github.com/dangerous-tech/MPGSMerchantServer /app
#Checkout staging if required
RUN if [ "${TRAVIS_BRANCH}" = "staging" ]; then git checkout staging; fi
RUN if [ "${BUILD_ENV}" = "arm" ]; then GOOS=linux GOARCH=arm go build .; fi
RUN if [ "${BUILD_ENV}" = "arm64" ]; then GOOS=linux GOARCH=arm64 go build .; fi
RUN if [ "${BUILD_ENV}" = "amd64" ]; then GOOS=linux GOARCH=amd64 go build .; fi
#Catch in case this is a separate build (i.e. not with buildkit)
RUN if [ ! -f mpgsmerchantserver ]; then go build .; fi

#Run app in thin container
FROM alpine:latest
WORKDIR /app
ENV PORT=8080
COPY --from=GO /app/mpgsmerchantserver .
ENTRYPOINT ["./mpgsmerchantserver"]
