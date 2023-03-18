FROM golang:latest

COPY [".", "/usr/src/"]

WORKDIR /usr/src/

RUN go install cmd/api-jokes.go

EXPOSE 8000

CMD ["api-jokes"]