FROM golang:1.14.3-alpine

RUN apk --no-cache update \
	&& apk --no-cache upgrade;

COPY ./main.go /app/main.go
CMD ["go", "run", "/app/main.go"]
