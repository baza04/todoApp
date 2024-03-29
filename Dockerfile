FROM golang:1.18-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download
RUN CGO_ENABLED=0 go build -o /todo-app ./cmd/web/main.go


FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /todo-app /todo-app
COPY ./configs/ /configs/
COPY ./wait-for-postgres.sh ./
COPY .env .

# install psql and make wait-for-postgres.sh executable
RUN apk --no-cache add postgresql-client && chmod +x wait-for-postgres.sh

CMD ["/todo-app"]