# Sample Blog Server
A simple blog server where articles can be created, searched for using an ID, and retrieved.

By default, the server is listening on port 8080.
This can be changed in the "config.json" and "docker-compose.yml" before running the server.

View server endpoints and details in the "api.raml" file.

# Features
- Create article
- Find article by ID
- List all articles


# Requirements
- Docker (and Docker Compose)
_For **macos user**, please use [Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)._
- golang v1.14 (optional unless intended to run tests)


# How to use
1. Download this [repository](https://github.com/jialijelly/sample_blog_server.git).
2. Please ensure "start.sh" is executable.
3. (a) Run "start.sh run" on command terminal. It set up the database and the blog server. OR
   (b) Run "start.sh test" on command terminal. It set up the database and run tests in the repository. 