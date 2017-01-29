# Gulp v3.9 with docker v1.13.0
# docker run --rm pottava/gulp:3.9-docker docker version
# docker run --rm pottava/gulp:3.9-docker docker-compose version

FROM pottava/gulp:3.9

ENV DOCKER_VERSION=1.13.0

RUN apk --no-cache add --virtual build-deps curl \

    # Install Docker
    && curl --location --silent --show-error -O \
        https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION}.tgz \
    && tar xzf docker-${DOCKER_VERSION}.tgz \
    && cp -r docker/* /usr/bin/ \

    # Clean up
    && rm -rf /usr/bin/completion \
    && apk del --purge -r build-deps
