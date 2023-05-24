# link_shortener
gRPC service for saving long links in database and getting their shortened versions
(make sure you have docker, golang installed)

**Endpoints:**
  1. SaveOriginal - сохраняет оригинальную (длинную) ссылку в базе данных и возвращает укороченный аналог
  2. GetOriginal - принимает укороченный аналог ссылки и возвращает длинную (оригинальную) её версию из базы данных

**Выбор базы данных**, выбор файла конфигурации и пароль для базы данных происходит в файле _.env_

**Running service**
(с помощью Makefile):
  1. make code-gen - загружает зависимости (через пакетный менеджер для linux), генерирует код сервиса из api/shortener.proto
  2. make run-server - запускает сервер (при запуске вне docker CONFIGFILE="configs/config.debug.json" в .env)
  3. make run-client - запускает клиент
  4. make docker - создает и запускает контейнеры с приложением и базой данных (при запуске в docker CONFIGFILE="configs/config.docker.json" в _.env_)
  5. make migration-up - создает таблицу для хранения ссылок в Postgres
  6. make migration-down - удаляет созданную таблицу
