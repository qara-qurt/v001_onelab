# v001_onelab
Чистая архитектура на практике

## Перед тем как запустить  `go mod download`  
## Build проетка - `docker build -t app .`
## Запустить проект -  `make run(запускает docker container)`  
## Остановить -  `make stop`

## Домашнее задание к уроку номер 2

## ДЗ:
Написать свой сервис хранения пользователей (фио, логин, пасс )
Реализовать graceful shudown
Чтение конфигов из ENV либо стандартные значения
Данные нужно хранить как in memory
## Бонус:
Хранение логов входящих запросов
Реализовать возможность увидеть процесс выполнения запроса в логах от начала и до сохранения в ДБ
Настроить линтер
## Что почитать:

Чистая Архитектура
системные сигналы linux и чем они отличаются
Виды переменных окружения и чем они отличаются
12 factor app
