# redirect

Tiny Go webserver that redirect to https

## Usage


```yml
version: '3.9'

services:
  redir:
    image: soliujing/go-redir-https:latest
    container_name: redir
    ports:
      - "80:80"
    networks:
      - bridge-all
    restart: unless-stopped
```
