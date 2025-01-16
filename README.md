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

- https://paniccaaa.github.io/crypto-observer/redoc.html
- https://paniccaaa.github.io/crypto-observer/swagger.html