# Golang Messenger

Небольшой backend pet-project на golang (мессенджер)


## Возможности

### Пользователи

- регистрация
- авторизация по JWT
- получение собственного профиля
- получение профиля другого пользователя
- изменение профиля
- удаление аккаунта

### Чаты

- создание приватных чатов
- получение списка чатов
- получение информации о чате
- удаление чата

### Сообщения

- отправка сообщений
- получение истории сообщений
- редактирование сообщений
- удаление сообщений

### WebSocket

- одно постоянное соединение на пользователя
- доставка новых сообщений в режиме реального времени
- поддержка нескольких устройств одновременно

---

## Стек

- Go
- PostgreSQL
- Gorilla WebSocket
- Docker Compose
- JWT
- Zap Logger

---

## Архитектура

Проект построен по feature-first структуре.

```
cmd/
internal/
    core/
    feature/
        users/
        chats/
        messages/
migrations/
```

Каждая feature содержит:

- repository
- service
- transport

---

## Запуск

### 1. Клонировать проект

```bash
git clone ...
cd golang-messenger
```

### 2. Создать .env

```
cp .env.example .env
```

Заполнить значения:

```
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
DATABASE_URL=
JWT_SECRET=
```

JWT_SECRET можно сгенерировать:

```bash
openssl rand -base64 32
```

---

### 3. Запустить PostgreSQL

```
make env-up
```

---

### 4. Применить миграции

```
make migrate-up
```

---

### 5. Запустить порт-форвардинг

```
make env-port-forward
```

---

### 6. Запустить сервер

```
make messenger-run
```

---

## Переменные окружения

|      Переменная       |           Описание            |
|-----------------------|-------------------------------|
| HTTP_ADDR             | адрес HTTP сервера            |
| HTTP_SHUTDOWN_TIMEOUT | graceful shutdown timeout     |
| JWT_SECRET            | секрет JWT                    |
| JWT_TTL               | время жизни JWT               |
| DATABASE_URL          | строка подключения PostgreSQL |
| LOGGER_LEVEL          | уровень логирования           |
| LOGGER_FOLDER         | папка логов                   |
| TZ                    | часовой пояс                  |

---

## Make команды

```
make env-up
make env-down
make migrate-up
make migrate-down
make migrate-create seq=init
make messenger-run
```

---

## Что планируется

- [ ] Redis
- [ ] Kafka / RabbitMQ
- [ ] загрузка файлов
- [ ] групповые чаты
- [ ] уведомления
- [ ] Docker образ приложения
- [ ] CI/CD

## О проекте

Цель проекта — изучить разработку production-подобного backend-приложения на Go.

При разработке особое внимание уделяется:

- разделению ответственности между слоями;
- расширяемой архитектуре;
- работе с PostgreSQL;
- WebSocket-соединениям;
- инфраструктуре приложения;
- читаемости кода.