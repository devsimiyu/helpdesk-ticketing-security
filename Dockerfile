##
## Build
##

# use alpine for its small footprint compared to Ubuntu
FROM golang:1.18-alpine as BUILD

# set working directory
WORKDIR /app

# install module dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy source code files
COPY *.go ./

# compile the app
RUN go build -o server


##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

# Copy artifacts from build stage
COPY --from=BUILD /server /server


EXPOSE 8080
USER nonroot:noroot

# launch the server
ENTRYPOINT ["/server"]