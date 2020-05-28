FROM alpine:3.10

RUN apk --no-cache update \
	&& apk --no-cache upgrade;

COPY ./runner.sh /app/run.sh
CMD ["/app/run.sh"]
