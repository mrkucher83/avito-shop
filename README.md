# **Тестовое задание для стажёра Backend-направления (зимняя волна 2025)**

## Магазин мерча

## Запуск сервиса

- в корень проекта добавить .env файл со следующими переменными окружения:
  - DB_HOST
  - DB_PORT
  - DB_USER
  - DB_PASSWORD
  - DB_NAME
  - SERVER_PORT
  - AVITO_SECRET
  - MAC_OS_USER or LINUX_USER
- миграции создаются при запуске сервиса (/pkg/helpers/pg/pg.go)
- docker-compose up --build
