# Person Service

Микросервис для управления пользователями.

## Запуск с Docker Compose

Для запуска всего стека (приложение + PostgreSQL) одной командой:

```bash
docker compose up
```

Для запуска в фоновом режиме:

```bash
docker compose up -d
```

## Переменные окружения

Приложение использует следующие переменные окружения:

- `POSTGRES_HOST` - хост PostgreSQL (по умолчанию: `localhost`)
- `POSTGRES_PORT` - порт PostgreSQL (по умолчанию: `5432`)
- `POSTGRES_USER` - пользователь PostgreSQL (по умолчанию: `person_user`)
- `POSTGRES_PASSWORD` - пароль PostgreSQL (по умолчанию: `password`)
- `POSTGRES_DB` - имя базы данных (по умолчанию: `person_service`)
- `HTTP_PORT` - порт HTTP сервера (по умолчанию: `8080`)

## API Endpoints

- `GET /person` - получить список пользователей
- `GET /person/:id` - получить пользователя по ID
- `POST /person` - создать пользователя
- `PUT /person/:id` - обновить пользователя
- `DELETE /person/:id` - удалить пользователя

## Примеры запросов

### Создание пользователя

```bash
curl -X POST http://localhost:8080/person \
  -H "Content-Type: application/json" \
  -d '{
    "email": "ivan.petrov@example.com",
    "phone": "123456789",
    "firstName": "Ivan",
    "lastName": "Petrov"
  }'
```

### Получение списка пользователей

```bash
curl http://localhost:8080/person
```

### Получение пользователя по ID

```bash
curl http://localhost:8080/person/1
```

## Остановка

```bash
docker compose down
```

Для удаления всех данных (включая volume с базой данных):

```bash
docker compose down -v
```

