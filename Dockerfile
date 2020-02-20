FROM alpine:3

RUN mkdir /app
WORKDIR /app
COPY ./bin/app .

RUN addgroup -g 1001 worker && \
    adduser --system --uid 1001 worker worker && \
    chown -R worker:worker /app

EXPOSE 8080
USER worker

CMD ./app
