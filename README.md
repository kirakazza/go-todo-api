# Todo API

Простое REST API для управления задачами с регистрацией и аутентификацией пользователей на Go с использованием Gin.

---

## О проекте

Это backend-приложение позволяет пользователям регистрироваться, логиниться и управлять своими todo-задачами (создавать, читать, обновлять, удалять).  
В проекте реализована аутентификация через JWT, все задачи привязаны к конкретному пользователю.

---

## Технологии

- Go 1.23  
- Gin (HTTP-фреймворк)  
- PostgreSQL (база данных)  
- JWT для аутентификации  
- Docker и docker-compose для удобного запуска и деплоя  

---

## Как запустить проект

### 1. Клонировать репозиторий

```bash
git clone https://github.com/yourusername/todo-api.git
cd todo-api
````

### 2. Создать `.env` файл

Создай файл `.env` в корне проекта со следующим содержимым (пример):

```env
PORT=8080
DB_USER=postgres
DB_PASSWORD=secretpassword
DB_NAME=todo_db
DB_HOST=db
DB_PORT=5432
JWT_SECRET=your_jwt_secret_key
```

### 3. Запустить PostgreSQL и сервер через Docker Compose

```bash
docker-compose up --build
```

Docker поднимет контейнер с базой данных и сервер.

### 4. Проверить работу

* Сервер слушает порт 8080
* Доступен эндпоинт `/ping` для проверки:

```bash
curl http://localhost:8080/ping
```

---

## Эндпоинты API

### Аутентификация

* `POST /register` — регистрация пользователя
  Тело запроса (JSON):

  ```json
  {
    "username": "username",
    "password": "password"
  }
  ```

* `POST /login` — логин пользователя
  Возвращает JWT токен при успешной аутентификации.

### Задачи (требуется авторизация через JWT в заголовке `Authorization: Bearer <token>`)

* `GET /todos` — получить список своих задач
* `POST /todos` — создать новую задачу
* `PUT /todos/:id` — обновить задачу (можно обновлять поля частично)
* `DELETE /todos/:id` — удалить задачу

---

## Пример запроса с JWT

```bash
curl -X GET http://localhost:8080/todos \
  -H "Authorization: Bearer <your_jwt_token>"
```

---

## Разработка

* Код разделён по папкам:

  * `internal/handler` — HTTP обработчики
  * `internal/service` — бизнес-логика
  * `internal/repository` — работа с базой данных
* Мидлвары и JWT аутентификация реализованы через Gin middleware

