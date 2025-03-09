#!/bin/sh
set -e

timeout=30
echo "Ожидание доступности базы данных (до $timeout секунд)..."
while ! pg_isready -h db -p 5432 -U "$DB_USER"; do
  sleep 2
  timeout=$((timeout - 2))
  if [ $timeout -le 0 ]; then
    echo "Ошибка: База данных недоступна!"
    exit 1
  fi
done

echo "Генерация SQL-кода..."
sqlc generate


echo "Применение миграций..."
/usr/local/bin/goose -dir sql/schema postgres "postgres://${DB_USER}:${DB_PASSWORD}@db:5432/${DB_NAME}?sslmode=disable" up
