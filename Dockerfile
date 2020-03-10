#Clone and build app
FROM golang:1.13.8-alpine
WORKDIR /app
ARG TRAVIS_BRANCH=$TRAVIS_BRANCH
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV PORT=8080
RUN apk --update add git
RUN git clone https://github.com/dangerous-tech/MPGSMerchantServer /app
#Checkout staging if required
RUN if [ "${TRAVIS_BRANCH}" = "staging" ]; then git checkout staging; fi
RUN go build -v .

#Run app in thin container
FROM alpine:latest
WORKDIR /app
ENV PORT=8080
COPY --from=GO /app/mpgsmerchantserver .
ENTRYPOINT ["./mpgsmerchantserver"]