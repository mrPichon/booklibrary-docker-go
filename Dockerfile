# syntax=docker/dockerfile:1

FROM golang:latest AS build
WORKDIR /app/
COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o booklibrary

FROM alpine:latest
WORKDIR /app/
# RUN apk add --no-cache bash
COPY --from=build /app/booklibrary .
COPY --from=build /app/static/* ./static/
COPY --from=build /app/templates/* ./templates/
COPY --from=build /app/fluentd/* ./fluentd/
# COPY --from=build /app/app.env .
CMD [ "./booklibrary" ]