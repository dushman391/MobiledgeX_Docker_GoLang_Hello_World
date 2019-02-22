## Creating a Simple Go Web Server

We'll create a simple web application in Go for demonstration in this article. This application will:
* Expose routes for the index route,
* Serve up a basic HTML page for /, and
* Use a configuration file to customize the application.

## Step 1 -- Create the Necessary Files
Upon completion, there will be three files within your temporary folder and the directory structure will look like:
 - Dockerfile -- Docker configurations go in here
 - index.html -- HTML page just to show that your Docker container works and can serve up a webpage
 - main.go -- GoLang app that serves up index.html

### Dockerfile

### index.html

### main.go

## Step 2 -- Build the Docker Image
To build the image, open a terminal, navigate to your temporary folder and type the following command

```
docker build . -t hello:1
```

## Step 3 -- Run the Image
To run the image, type the following command in the terminal window from the previous step:
``` 
docker run -it -p 30300:8001 hello:1
```
If this command is successful, you should immediately see a line of text that says "Server Started..."

## Step 4 -- Verify Server is Running
Open a browser to localhost:3030 and you should see the content of the index.html page displayed