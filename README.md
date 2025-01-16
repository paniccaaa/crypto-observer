# crypto-observer

1) Docker
```sh
docker compose build
docker compose up -d
```

- other way:
```sh
CONFIG_PATH="./config/dev.yaml" go run cmd/main.go 
```
2) Migrations
- Install -> https://github.com/pressly/goose
```sh
make goose-up
```

3) Docs:

gh-pages:
- https://paniccaaa.github.io/crypto-observer/public/redoc.html
- https://paniccaaa.github.io/crypto-observer/public/swagger.html

local:
- http://localhost:8089/docs/redoc.html
- http://localhost:8089/docs/swagger.html