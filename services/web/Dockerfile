FROM golang
ARG SERVICE_ENV
ENV SERVICE_ENV ${SERVICE_ENV:-"production"}
ENV PORT 8080
WORKDIR /go
ADD . /go/src/mono
RUN go install "mono/services/site"
RUN if [ ${SERVICE_ENV} = development ]; \
  then \
    echo "Development mode"; \
    go get github.com/pilu/fresh; \
  else \
    echo "Production mode"; \
  fi

WORKDIR src/mono/services/${SERVICE_NAME}

CMD if [ ${SERVICE_ENV} = development ]; \
  then \
    "fresh"; \
  else \
    "/go/bin/${SERVICE_NAME}"; \
  fi

EXPOSE ${PORT}
