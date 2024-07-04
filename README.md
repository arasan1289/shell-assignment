Steps to Run this application:
1. Clone the repository
2. Run go mod download
3. Create a file `config/config.yaml` in applications root directory. Example contents of the file is 
`
type: yml
name: Shell Assignment
env: development
port: 8080
url: 127.0.0.1
allowed_origins: 
- http://127.0.0.1
`