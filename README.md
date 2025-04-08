# TODO-лист


## Технологии
- Fiber
- PostgreSQL (pgx)
- Docker
- godotenv (парсинг .env)


## Запуск c помощью Docker
В файле `.env` уставновите `host.docker.internal` как значение переменной `POSTGRES_URL`:
```
POSTGRES_URL=host.docker.internal
```
Затем выполните команду:
```
docker compose up
```
## Локальный запуск
В файле `.env` уставновите `localhost` как значение переменной `POSTGRES_URL`:
```
POSTGRES_URL=localhost
```
Затем выполните команду:
```
go run .
```

## Переменные .env
- POSTGRES_USER - имя Postgres пользователя
- POSTGRES_PASSWORD - пароль пользователя
- POSTGRES_URL - адрес БД
- POSTGRES_PORT - порт БД
- POSTGRES_DB - назавание БД
- CREATE_DB - если нужно создать БД установить значение `true`
- PORT - порт API

![postman](https://ibb.co/xtNcq0xy)
![database](https://ibb.co/TDSJdSdT)