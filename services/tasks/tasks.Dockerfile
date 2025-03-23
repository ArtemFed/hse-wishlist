FROM golang:1.24.1 as build
WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN apt-get update && \
#    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go build -o tasks-svc ./services/tasks/cmd

FROM alpine:latest as production

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

COPY --from=build /app/tasks-svc ./

COPY --from=build /app/${ENV_FILE} ./services/tasks/.env
COPY --from=build /app/services/tasks/migrations ./services/tasks/migrations
COPY --from=build /app/services/tasks/config/config.docker.yml ./services/tasks/config/config.local.yml

CMD ["./tasks-svc"]

EXPOSE 8082
