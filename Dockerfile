FROM golang:1.12.6-alpine3.10 AS build

RUN apk add --no-cache git


WORKDIR $GOPATH/src/github.com/mchirico/gomini

# Copy the entire project and build it

COPY . $GOPATH/src/github.com/mchirico/gomini
COPY test-fixtures/data.csv /data.csv
RUN go get -v  -d .
RUN go build -o /bin/project

# This results in a single layer image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /bin/project /bin/project
COPY --from=build /data.csv /data.csv
ENTRYPOINT ["/bin/project"]
CMD ["dataroot"]

