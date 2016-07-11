# Gulp v3.9
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 npm install
# docker run --rm -v `pwd`:/app pottava/gulp:3.9 gulp

FROM pottava/nodejs:6.3

RUN npm install --global gulp@3.9.1

WORKDIR /app
VOLUME /app

CMD ["gulp"]
