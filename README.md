Steps to Run this application:
1. Clone the repository
2. Run go mod download
3. Create a file `config/config.yaml` in application root directory. Example contents of this file is 
```
type: yml
name: Shell Assignment
env: development
port: 8080
url: 127.0.0.1
allowed_origins: 
- http://127.0.0.1
```
4. Build a docker image by running `docker build -t <container-name> .`
5. Run the container by running `docker run -p 8080:8080 <container-name>`
