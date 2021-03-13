# TODO
To build for EdgeRouter X:
```
GOOS=linux GOARCH=mipsle go build -o edge
```

# Environment variables
You can place them in an .env file
```
CF_API_EMAIL=''
CF_API_KEY=''
CF_ZONE=''
CF_RECORD_NAME=''
PUSHOVER_TOKEN=''
PUSHOVER_USER=''
IFACE_NAME='pppoe0'
CGNAT_RANGE='100.64.0.0/10'
```