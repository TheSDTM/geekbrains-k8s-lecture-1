# О проекте
Учебное приложения для проведения первой лекции по интенсиву Kubernetes от GeekBrains.

В папке Kubernetes находятся все конфигурационные файлы, которые используются во время лекции.


# Полезные команды
Ниже приведен список команд, которые могут понадобиться во время лекции.

## Установка кластера
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
