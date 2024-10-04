FROM golang:1.22.7 as build
WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN apt-get update && \
#    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go build -o tasks-svc ./services/tasks/cmd

FROM alpine:latest as production

COPY --from=build /app/tasks-svc ./

COPY --from=build /app/.env ./services/tasks/
COPY --from=build /app/services/tasks/migrations ./services/tasks/migrations
COPY --from=build /app/services/tasks/config/config.docker.yml ./services/tasks/config/config.local.yml

CMD ["./tasks-svc"]

EXPOSE 8082