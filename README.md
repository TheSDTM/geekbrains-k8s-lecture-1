# О проекте
Учебное приложения для проведения первой лекции и выполнению ДЗ по интенсиву Kubernetes от GeekBrains.

В папке Kubernetes находятся все конфигурационные файлы, которые используются во время лекции.

# Что нужно выполнить до начала первой лекции
1. установить [kubectl](https://kubernetes.io/ru/docs/tasks/tools/install-kubectl/) для работы с кластером через командную строку
2. установить [Kubenav](https://github.com/kubenav/kubenav/releases) для работы с кластером через GUI
3. установить [minikube удобным способом](https://kubernetes.io/ru/docs/tasks/tools/install-minikube/). ВНИМАНИЕ! Если у вас Windows, то не используйте связку WLS (Windows  Linux Subsystem) и Docker, так как это приведет к проблемам в в работе Kubernetes. Рекомендуется использовать HyperKit.
4. создать новый кластер (см. команды ниже в разделе "полезные команды")
5. установить ingress controller (см. команды ниже в разделе "полезные команды")
6. убедиться в работоспособности кластера командой `kubectl get ns`
7. клонировать этот репозиторий себе на компьютер
8. ознакомиться с файлом README.md

# Что нужно выполнить до начала второй лекции
1. установить [helm](https://helm.sh/) для работы с чартами (пакетами)

# Домашнее задание
## Описание задачи
Подготовить yaml файлы для развертывания приложения (сделано на лекции), развернуть его в своём кластере и убедиться в доступности и корректной работе. Корректная работа приложения означает, что заголовки на карточках горят зелеными. Ниже приведены требования к развертываниям отдельных компонентов системы.

### Backend
В конфигурации Deployment предусмотреть:
- количество реплик должно быть больше одной
- метки (для конфигурации Service)
- liveness probe
- startup prob
- аффинность по hostname

В конфигурации Service предусмотреть:
- сервис верно указывает на Pod’ы 
- должен быть типа ClusterIP

### Frontend
В конфигурации Deployment предусмотреть:
- количество реплик должно быть больше одной
- метки (для конфигурации Service)
- liveness probe
- startup probe
- аффинность по hostname

В конфигурации Service предусмотреть:
- сервис верно указывает на Pod’ы 
- должен быть типа ClusterIP

### Redis, RabbitMQ, PostgreSQL
В конфигурации Deployment предусмотреть:
- должна быть одна реплика
- метки (для конфигурации Service)

В конфигурации Service предусмотреть:
- сервис верно указывает на Pod’ы 
- должен быть типа ClusterIP

Для базы данных PostgreSQL обеспечить постоянное хранилище с использованием PersistentVolumeClaim.

### Общее
Все должно быть развернуто в namespace названием “k8s-course”.
Также нужно подготовить конфиг Ingress в соответствии с описанием работы приложения (см. репозиторий).

## Что нужно отправить на проверку
Набор YAML файлов, которыми вы создали все необходимые сущности Kubernetes.

# Состав проекта
Папка app содержит два приложения - frontend и backend.

## Backend
Представляет собой простое API на порту 80, написанное на Go, с двумя методами:
1. GET / - печать на экран информации о подключении к БД, доступности секретов и т.д.
2. GET /api - возврат той же информации, но в формате JSON

Backend имитирует взаимодействие с:
1. секретом - считывание и вывод данных из секрета, который должен находиться по пути `/etc/geekbrains/username`
2. файлом конфига - считывание и вывод данных из конфига, который должен находиться по пути `/config/config.yaml`
3. переменными окружения - считывание и вывод переменной `SOME_VARIABLE`
4. постоянным хранилищем - запись и вывод данных из файла `/data/data`
5. PostgreSQL - проверка подключения и выполнение запроса
6. Redis - проверка подключения и пинг
7. RabbitMQ - проверка подключения и пинг

### Docker
Контейнер backend называется `yaroslavperf/gb-k8s-lecture-1` и имеет 2 тега: `latest` и `1.0`.
Ссылки на DockerHub для зависимостей:
|Название|Ссылка|
|-|-|
|RabbitMQ|[перейти](https://hub.docker.com/_/rabbitmq)|
|Redis|[перейти](https://hub.docker.com/_/redis)|
|PostgreSQL|[перейти](https://hub.docker.com/_/postgres)|

### Переменные окружения
|Название|Описание|Пример|Примечание|
|-|-|-|-|
|`RABBIT_HOST`|Адрес хоста RabbitMQ|127.0.0.1:5672||
|`RABBIT_USER`|Имя пользователя RabbitMQ|guest|В RabbitMQ пользователь имеет логин по-умолчанию guest|
|`RABBIT_PASSWORD`|Пароль пользователя RabbitMQ|guest|В RabbitMQ пользователь имеет пароль по-умолчанию guest|
|`REDIS_HOST`|Адрес хоста Redis|127.0.0.1:6379||
|`REDIS_PASSWORD`|Пароль Redis|abc123|По-умолчанию в контейнере Redis не установлен пароль. Т.е. не нужно его указывать в этой переменной окружения, если настройки Redis по-умолчанию не менялись|
|`POSTGRE_ADDR`|Адрес хоста PostgreSQL|127.0.0.1:5432||
|`POSTGRE_DB`|Название БД PostgreSQL|db|Эта переменная должна иметь то же значение, что и переменная окружения `POSTGRES_PASSWORD` для контейнера PostgreSQL|
|`POSTGRE_USER`|Логин пользователя PostgreSQL|user|Эта переменная должна иметь то же значение, что и переменная окружения `POSTGRES_USER` для контейнера PostgreSQL|
|`POSTGRE_PASSWORD`|Пароль пользователя PostgreSQL|password|Эта переменная должна иметь то же значение, что и переменная окружения `POSTGRES_PASSWORD` для контейнера PostgreSQL|
|`SOME_VARIABLE`|Переменная, которая будет выведена на экран|hello, world|Конкретные значения никак не влияет на работу|

## Frontend
Приложение на порту 80, которое только лишь обращается к методу `/api` backend и визуализирует ответ. В случае, если тест прошел успешно, выводится зеленый заголовок соответствующего блока. Если тест не прошел - выводится красный заголовок и сообщение об ошибке.

### Docker
Контейнер frontend называется `yaroslavperf/gb-k8s-frontend-lecture-1` и имеет 2 тега: `latest` и `1.0`.

### Обращение к backend
Frontend ожидает, что backend находится по тому же адресу, что frontend, но по другому пути - /api.
То есть структура выглядит так:
- GET / - запросы уходит на frontend
- GET /api - запросы уходят на backend

Достигнуть такого результата можно за счет использования двух путей в одном Ingress.

# Полезные команды
Ниже приведен список команд, которые могут понадобиться во время лекции.

## Создание кластера
minikube start --nodes 2 --vm=true  
minikube addons enable ingress

## Удаление кластера
minikube delete

## Остановка кластера
minikube stop

## Включение кластера
minikube start

## Применить YAML конфигурацию
kubectl apply -f 00-file.yaml

## Удалить YAML конфигурацию
kubectl delete -f file.yaml

## Получить доступ к приложению через сервис
minikube service myservice -n test

## Получить IP-адрес Ingress
kubectl get ingress -n test

## Установка cert manager
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml

## Установка Prometheus Operator
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus-operator prometheus-community/kube-prometheus-stack -n test-2 --values=prometheus-values.yaml
```
