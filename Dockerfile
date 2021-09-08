FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/ip-resolver
ARG IP_RESOLVER_SERVER_PORT

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++ git curl lsof
RUN go mod tidy

RUN cd $GOPATH/src/ip-resolver/ && go build main.go
RUN chmod 0766 $GOPATH/src/ip-resolver/scripts/init.sh

# This application is exposed through IP_RESOLVER_SERVER_PORT to the outside
# See .env to change the value.
EXPOSE $IP_RESOLVER_SERVER_PORT

# Run the executable
CMD ["./scripts/init.sh"]