FROM alpine:edge

WORKDIR /app

RUN apk --no-cache add tzdata

ADD invoke /app/

EXPOSE 9000

CMD [ "./invoke" ]
