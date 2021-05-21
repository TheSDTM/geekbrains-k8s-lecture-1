# О проекте
Учебное приложения для проведения первой лекции по интенсиву Kubernetes от GeekBrains.

В папке Kubernetes находятся все конфигурационные файлы, которые используются во время лекции.

# Что нужно выполнить до начала лекции
1. установить [kubectl](https://kubernetes.io/ru/docs/tasks/tools/install-kubectl/) для работы с кластером через командную строку
2. установить [Kubenav](https://github.com/kubenav/kubenav/releases) для работы с кластером через GUI
3. установить [minikube удобным способом](https://kubernetes.io/ru/docs/tasks/tools/install-minikube/)
4. создать новый кластер (см. команды ниже в разделе "полезные команды")
5. убедиться в работоспособности кластера командой `kubectl get ns`
6. клонировать этот репозиторий себе на компьютер

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
