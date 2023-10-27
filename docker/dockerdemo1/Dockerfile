#
# Dockerfile to create lightweight go-executables by means of staged dockercontainer
#
# Prerequisite: A volume needs to be attached in Docker
#
#
#
# Use the image golang:alpine from dockerhub to support golang build on alpine linux
# 
FROM golang:alpine
#

# Copy all files from the source-directory
COPY . /./
#
# Set the actual WORKDIR to the build-directory
WORKDIR /./
#
# Download Go modules and dependencies
RUN go mod download
#
# Build the executable and name it main
RUN go build -o ./main .
#
# Start the executable main in the actual WORKDIR (/app)
CMD ["./main"]
#
###### EOF ######
