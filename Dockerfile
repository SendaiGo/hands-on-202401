FROM        golang:1.21.0 AS build-env
ARG         BUILD_ENV="codebuild"
RUN         if [ ${BUILD_ENV} = "local" ]; then \
              go install github.com/go-delve/delve/cmd/dlv@v1.22.0; \
              go install github.com/cosmtrek/air@latest; \
            fi
ENV         WORKDIR_PATH /go/src/github.com/sendaigo/hands-on-202401
WORKDIR     ${WORKDIR_PATH}
ADD         . ${WORKDIR_PATH}
RUN         CGO_ENABLED=0 go build -buildvcs=false -o /bin/hands-on-202401 ./cmd/hands-on-202401

FROM        scratch
COPY        --from=build-env /bin/hands-on-202401 /bin/hands-on-202401
EXPOSE      80
ENTRYPOINT [ "/bin/hands-on-202401" ]