# Start from base image 1.16.4
FROM golang:1.16.4

# Configure the repo url so we can configure our work directory
ENV REPO_URL=github.com/katsun0921/bookstore_items-api

# Setup out GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPR_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o items-api .

EXPOSE 8081

CMD ["./items-api"]
