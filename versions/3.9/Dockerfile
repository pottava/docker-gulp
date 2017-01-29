# Gulp v3.9
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 npm install
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 gulp

FROM alpine:3.5

ENV GULP_VERSION=3.9.1 \
    NODE_VERSION=7.2.1-r1

RUN apk --no-cache add tini ca-certificates nodejs-current=${NODE_VERSION} \

    # Install Gulp
    && npm install --global "gulp@${GULP_VERSION}" \

    # Clean up
    && rm -rf /root/.npm \
    && find / -depth -type d -name test* -exec rm -rf {} \; \
    && find / -depth -type f -name *.md -exec rm -f {} \; \
    && find / -depth -type f -name *.yml -exec rm -f {} \;

WORKDIR /app
VOLUME /app

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["gulp"]
