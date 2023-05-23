# link_shortener
gRPC service for saving long links in database and getting their shortened versions
(make sure you have docker, golang installed)

**Endpoints:**
  1. SaveOriginal - сохраняет оригинальную (длинную) ссылку в базе данных и возвращает укороченный аналог
  2. GetOriginal - принимает укороченный аналог ссылки и возвращает длинную (оригинальную) её версию из базы данных

**Running service**
(с помощью Makefile):
  1. make code-gen - загружает зависимости (через пакетный менеджер для linux), генерирует код сервиса из api/shortener.proto
  2. make run-server - запускает сервер
  3. make run-client - запускает клиент
