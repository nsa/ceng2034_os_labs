FROM golang:alpine

WORKDIR /some_random/working_dir
ADD files /some_random/working_dir/files
ADD assignment3.go .
ADD tidyTidy.sh .
ADD entrypoint.sh .

RUN apk update && \
	apk add --no-cache bash git

RUN go get github.com/gabriel-vasile/mimetype

CMD ["sh", "./entrypoint.sh" ]