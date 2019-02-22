# MobiledgeX Docker GoLang Hello World 
# 
FROM golang:alpine

# set maintainer
LABEL maintainer "alexander.donn@mobiledgex.com"

# Copy the application files (needed for production)
ADD . /app

# Set the working directory to the app directory
WORKDIR /app

CMD ["go", "run", "main.go"]

# tell docker what port to expose
# This doesn't seem to do anything?
EXPOSE 8000