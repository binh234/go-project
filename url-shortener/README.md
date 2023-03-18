# URL shortener with Go and Redis

This repo is an implementation of a URL shortener service using Golang, Fiber, and Redis. Each user can send a maximum of 10 requests per 30 minutes to shorten their URLs. Every shortened URL will have an expiration of 24 hours by default. For customization, users can suggest  and set their own shortened URL as well as the expiration time (up to 720 hours).

To continue using the shortened URL, the user also can extend the old expiration using the same URL and the old shortened URL with the new expiration time.

## Features

- [x] Shorten URL
- [x] Customize short URL
- [x] Customize expiration
- [x] Extend expiration

## Setup

Create a `.env` file inside `api` folder similar to this:

```text
DB_ADDR="db:6379"
DB_PASS=""
APP_PORT=":3000"
DOMAIN="localhost:3000"
API_QUOTA=10
MAX_TRY=5
```

You can modify the `DB_ADDR` and `DB_PASS` to use your own Redis server

For deployment, modify the `DOMAIN` to your deployment URL.

## Start service with Docker Compose

### To start the application

Step 1: start Redis and Go app

```bash
docker-compose -f docker-compose.yaml up
```

#### Request and response struct

```Go
type request struct {
 URL         string        `json:"url"`
 CustomShort string        `json:"short"`
 Expiry      time.Duration `json:"expiry"`
}

type response struct {
 URL             string        `json:"url"`
 CustomShort     string        `json:"short"`
 Expiry          time.Duration `json:"expiry"`
 XRateRemaining  int           `json:"rate_limit"`
 XRateLimitReset time.Duration `json:"rate_limit_reset"`
}
```

#### Shorten URL

```text
    POST http://localhost:3000/api/v1
```

Example:

```json
{
    url: "https://www.youtube.com/watch?v=p-6pvjsH_ic"
}
```

#### Resolve URL

```text
    GET http://localhost:3000/{id}
```

### To build a docker image from the application

```bash
docker build -t mygo:1.0 ./api       
```
