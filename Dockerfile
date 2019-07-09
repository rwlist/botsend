FROM    golang:1.12.7-alpine3.9 
ENV     GO111MODULE=on
WORKDIR /app
RUN     apk update && apk add ca-certificates git
COPY    go.mod .
COPY    go.sum .
RUN     go mod download
COPY    . .
RUN     go build -v .
ENTRYPOINT [ "./botsend" ]