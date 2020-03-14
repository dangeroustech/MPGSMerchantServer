#!/bin/bash

set -ev
export DOCKER_CLI_EXPERIMENTAL=enabled

DOCKER_IMAGE="mpgsmerchantserver"
DOCKER_ORG="dangeroustech"

# Set Correct Branch
if [ "${TRAVIS_BRANCH}" = "master" ]; then
    DOCKER_TAG="stable"
else
    DOCKER_TAG="staging"
fi

# If This Isn't A PR, Push to Dockerhub
if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
    docker login -u $DOCKER_USER -p $DOCKER_PASSWORD

    docker manifest create ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG} \
            ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-amd64 \
            #${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-arm \
            #${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-arm64

    docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-amd64 --arch amd64
    #docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-arm --arch arm
    #docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}-arm64 --arch arm64

    docker manifest push ${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}

        docker manifest create ${DOCKER_ORG}/${DOCKER_IMAGE}:latest \
                ${DOCKER_ORG}/${DOCKER_IMAGE}:latest-amd64 \
                #${DOCKER_ORG}/${DOCKER_IMAGE}:latest-arm \
                #${DOCKER_ORG}/${DOCKER_IMAGE}:latest-arm64

        docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:latest ${DOCKER_ORG}/${DOCKER_IMAGE}:latest-amd64 --arch amd64
        #docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:latest ${DOCKER_ORG}/${DOCKER_IMAGE}:latest-arm --arch arm
        #docker manifest annotate ${DOCKER_ORG}/${DOCKER_IMAGE}:latest ${DOCKER_ORG}/${DOCKER_IMAGE}:latest-arm64 --arch arm64

        docker manifest push ${DOCKER_ORG}/${DOCKER_IMAGE}:latest
else
    # If This is a PR, Check Images and Tags Are Correct
    docker images
fi