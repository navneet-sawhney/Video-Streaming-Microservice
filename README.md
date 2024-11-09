**Step-by-Step Process for Building, Deploying, and Running the Video-Streaming Microservice**

1. Go Application Code (main.go)
In the main.go file, I implemented a basic HTTP server that serves the video file

Serving Video: The server listens on port 3000 and handles requests to the /video endpoint. When a request is received, it opens the video file (Vectors.mp4), streams it, and serves it to the client.
Headers: I set the appropriate headers like Content-Type: video/mp4 and Accept-Ranges: bytes to ensure that the video can be streamed properly.
Error Handling: I also included error handling for file opening and retrieving file metadata to make sure everything runs smoothly.


2. Dockerfile Explanation
I used official Go image (golang:1.23.1) as the base for building the video-streaming microservice. Here is how I structured the Dockerfile:

Base Image: I used the golang: 1.23 image, which provides the necessary Go environment for building the application.
Working Directory: I set the working directory to /app inside the container using the WORKDIR directive.
Copying Files: I copied the essential files (main.go, go.mod and Vectors.mp4) into the container. 
Building the Application: I ran the go build command to compile the Go program into an executable called video-streaming. 
Exposing Port: I exposed port 3000 so that the video-streaming service would be accessible on this port.
Running the Application: The CMD command runs the built application (../video-streaming) when the container starts.


3. Running the Docker Container
Build the Docker Image: I ran the following command to build the Docker image:

docker build -t video-streaming



Run the Container: Then, I ran the container with the command:

docker run -d -p 3000:3000 video-streaming

This command starts the container in detached mode (-d), mapping port 3000 of the container to port 3000 on my local machine.

Verify the Container is Running: I checked that the container was running with:

docker ps

This listed all the running containers, and I found the CONTAINER ID for the video-streaming container.

Check the Logs: To verify the server started successfully, I ran:

docker logs <CONTAINER_ID>


4. Testing the Video Stream
To test the video-streaming functionality, I:

I navigated to localhost:3000/videoLinks to an external site. in my browser.

Output- The video played directly in the browser, which meant the server was successfully streaming the video file.



*References*

Davis, Ashley. Bootstrapping Microservices with Docker, Kubernetes, and Terraform. Shelter Island, NY: Manning Publications, 2021.

Gajjar, Amit. "Build a Containerized Microservice in Golang." Velotio Engineering Blog, October 9, 2020. https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golangLinks to an external site..

 
