# GURL

Curl alternative written in Golang.

## Installation:

Clone the repo, run `go build` and add it to PATH.

## Usage:

### GET Request

```
GURL https://www.url.com -H "Accept: application/json"
```

### POST Request

```
GURL https://www.url.com -X POST -d "{'data':'data'}
```

### PUT Request

```
GURL https://www.url.com -X PUT -d "{'data':'data'}
```
