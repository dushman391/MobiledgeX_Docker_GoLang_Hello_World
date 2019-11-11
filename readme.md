## Creating a Simple Go Web Server

We'll create a simple web application in GoLang to demonstrate the process of "dockerizing" a GoLang app   in this article. This application will:
* Create a Dockerized GoLang webserver to serve up an HTML page
* Walk through the "Dockerfile" configuration file

## Step 1 -- Create the Necessary Folders & Files
Create a temporary folder with the name of your choosing. I like pizza, so I'm calling my folder "pizza1".

Next create the following three files within the folder and the directory structure should look like:
<pizza1>
 - Dockerfile -- Docker configurations go in here
 - index.html -- HTML page just to show that your Docker container works and can serve up a webpage
 - main.go -- GoLang app that serves up index.html

## Step 2 -- Build the Docker Image
To build the image, open a terminal, navigate to your temporary folder and type the following command

```
docker build . -t hello:1
```

## Step 3 -- Run the Image
To run the image, type the following command in the terminal window from the previous step:
``` 
docker run -it -p 3030:8001 hello:1
```
If this command is successful, you should immediately see a line of text that says "Server Started..."

## Step 4 -- Verify Server is Running
Open a browser to localhost:3030 and you should see the content of the index.html page displayed
