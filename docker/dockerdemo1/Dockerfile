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
COPY . ./import
#
# Set the actual WORKDIR to the directory where the app will be built
WORKDIR import
#
# Download Go modules and dependencies
RUN go mod init import
RUN go mod tidy
#
# Build the executable and name it app in root-dir
RUN go build -o ./app .
#
# Start the executable "app" in the actual WORKDIR (./)
CMD ["./app"]
#
###### EOF ######
