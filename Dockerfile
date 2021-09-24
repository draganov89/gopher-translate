FROM golang:1.16-alpine

WORKDIR /gopher-translate

COPY . .

# nothing to download for now
RUN go mod download

RUN go build -o ./gopher-translate

EXPOSE 1234

CMD [ "./gopher-translate", "--port", "1234" ]