# link_shortener
gRPC service for saving long links in database and getting their shortened versions
(make sure you have docker, golang installed)

# gRPC endpoints:
  1. SaveOriginal - сохраняет оригинальную (длинную) ссылку в базе данных и возвращает укороченный аналог
  2. GetOriginal - принимает укороченный аналог ссылки и возвращает длинную (оригинальную) её версию из базы данных

# REST endpoints:
  1. **POST** :8081/save - сохранение полной ссылки и получение укороченной версии
      ```
      {
        "originalLink": "[original link]"
      }
      ```
  2. **GET** :8081/get/{short_link} - получение полной ссылке по укороченной версии

**Выбор базы данных**, выбор файла конфигурации и пароль для базы данных происходит в файле _.env_

# Running service
(с помощью Makefile):
  1. _make code-gen_ - загружает зависимости (через пакетный менеджер для linux), генерирует код сервиса из api/shortener.proto
  2. _make run-server_ - запускает сервер (при запуске вне docker CONFIGFILE="configs/config.debug.json" в .env)
  3. _make run-client_ - запускает клиент
  4. _make docker_ - создает и запускает контейнеры с приложением и базой данных (при запуске в docker CONFIGFILE="configs/config.docker.json" в _.env_)
  5. _make migration-up_ - создает таблицу для хранения ссылок в Postgres
  6. _make migration-down_ - удаляет созданную таблицу
