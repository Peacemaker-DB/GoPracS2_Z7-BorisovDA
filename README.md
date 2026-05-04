# Практическое задание № 7 Борисов Денис Александрович ЭФМО-01-25
Тема: Написание Dockerfile и сборка контейнера

Задачи практики:
1.	Изучить назначение контейнеризации в backend-разработке. 
2.	Понять различие между Docker-образом и контейнером. 
3.	Освоить структуру и назначение файла Dockerfile. 
4.	Научиться использовать multi-stage build для сборки Go-приложения. 
5.	Освоить настройку. dockerignore. 
6.	Научиться собирать Docker-образ локально. 
7.	Научиться запускать контейнер с пробросом портов и передачей переменных окружения. 
8.	Освоить базовый запуск нескольких сервисов через Docker Compose. 
9.	Научиться проверять работу контейнеризированного приложения.

Выполнение практического задания.

1.	Структура проекта

<img width="361" height="398" alt="image" src="https://github.com/user-attachments/assets/7eff8914-111b-4f84-b09c-cbb60758d291" />

2.	Создание файла docker-compose.yml.

<img width="436" height="327" alt="1" src="https://github.com/user-attachments/assets/bb68e59d-0cb2-4b90-b857-80ad84923b7b" />


Созданный Dockerfile для app

<img width="872" height="505" alt="2" src="https://github.com/user-attachments/assets/94bcdd6d-3a48-41ec-8302-a7277fd250c9" />

Созданный dockerignore

<img width="332" height="240" alt="3" src="https://github.com/user-attachments/assets/52f0ee79-9f2e-4880-8434-8cbbb668fa7a" />

3. Тестирование

Запуск сервера

<img width="821" height="71" alt="5" src="https://github.com/user-attachments/assets/1fc2730b-0f43-4a06-b425-86464730a605" />

Развертывание сервисов

<img width="1025" height="181" alt="7" src="https://github.com/user-attachments/assets/ae807df4-8881-4909-a5db-7f230ae2471c" />

Работающий контейнер

<img width="1570" height="335" alt="8" src="https://github.com/user-attachments/assets/e3679ad2-d7e4-4690-86c5-25e2c0aec01c" />

Тест состояния

<img width="943" height="630" alt="9" src="https://github.com/user-attachments/assets/f97df871-452e-4850-b8a6-01fbb0432dd8" />
