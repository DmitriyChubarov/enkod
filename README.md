# Person Service

CRUD сервис на golang

## Запуск с Docker Compose

Для запуска одной командой:

```bash
docker compose up --build
```

## Переменные окружения

Необходимо создать .env файл:

```bash
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=person_user
POSTGRES_PASSWORD=password
POSTGRES_DB=person_service
HTTP_PORT=8080
```

## API Endpoints

- `GET /person` - получить список пользователей
- `GET /person/id` - получить пользователя по ID
- `POST /person` - создать пользователя
- `PUT /person/id` - обновить пользователя
- `DELETE /person/id` - удалить пользователя

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

### Контакты
- tg: @eeezz_z
- gh: https://github.com/DmitriyChubarov

