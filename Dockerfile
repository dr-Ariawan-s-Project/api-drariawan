FROM golang:1.19 AS build-stage

ARG APP_PATH="/app"
ENV APP_PATH=${APP_PATH}

ENV GO111MODULE=on

RUN mkdir ${APP_PATH}

WORKDIR ${APP_PATH}

COPY . .

RUN apt-get update \
    && export GOPROXY=https://proxy.golang.org \
    && go mod tidy -e

# build binary and put in PATH folder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o service-app ${APP_PATH}/main.go;

FROM debian:10.11

ARG TZ="Asia/Jakarta"
ARG APP_PATH="/app"
ARG APP_PORT=80

ENV APP_PATH=${APP_PATH}
ENV APP_PORT=${APP_PORT}
ENV TZ=${TZ}

RUN addgroup --system appgroup && adduser --system --disabled-password --no-create-home --home ${APP_PATH} --ingroup appgroup appuser

COPY --chown=appuser:appgroup --from=build-stage ${APP_PATH}/service-app ${APP_PATH}/
COPY --chown=appuser:appgroup --from=build-stage ${APP_PATH}/migration/ ${APP_PATH}/migration/
COPY --chown=appuser:appgroup --from=build-stage ${APP_PATH}/utils/files/ ${APP_PATH}/utils/files/

RUN apt-get update && apt-get install -y iputils-ping telnet wget

# setup TZ
RUN apt-get install -y tzdata; \
  cp /usr/share/zoneinfo/${TZ} /etc/localtime; \
  update-ca-certificates 2>/dev/null || true;

# go migrate tool
# ref: https://github.com/golang-migrate/migrate
ENV GOMIGRATE_VERSION v4.15.2
RUN wget https://github.com/golang-migrate/migrate/releases/download/$GOMIGRATE_VERSION/migrate.linux-amd64.tar.gz
RUN tar -C /usr/local/bin -xzvf migrate.linux-amd64.tar.gz
RUN ls -a usr/local/bin/migrate
RUN rm migrate.linux-amd64.tar.gz

USER appuser

WORKDIR ${APP_PATH}

RUN mkdir export
RUN mkdir public

STOPSIGNAL SIGINT

# make sure this port same as your app port
EXPOSE ${APP_PORT}

CMD [ "./service-app" ]