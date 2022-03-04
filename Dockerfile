FROM golang:alpine
WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o fileserver . && mkdir tmp

CMD ["./fileserver"]
# RUN mv fileserver /usr/bin/fileserver
# ENTRYPOINT ["fileserver"]   