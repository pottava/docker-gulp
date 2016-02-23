# Gulp v3.9
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 npm install
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 gulp

FROM pottava/nodejs:stable

RUN npm install --global gulp

WORKDIR /app
VOLUME /app

CMD ["gulp"]
