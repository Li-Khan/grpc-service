Домашнее задание - GRPC сервис
==============================
Цель
----
Создать GRPC API для сервиса календаря Тех. задание: 
https://github.com/OtusTeam/Go/blob/master/project-calendar.md 

Цель данного занятия: отработка навыков работы с GRPC, построение современного API.
- Создать отдельную директорию для Protobuf спек.
- Создать Protobuf спеки с описанием всех методов API, их объектов запросов и ответов.
- Т.к. объект Event будет использоваться во многих ответах разумно выделить его в отдельный message.
- Создать отдельный директорию для кода GRPC сервера
- Сгенерировать код GRPC сервера на основе Protobuf спек (скрипт генерации сохранить в репозиторий).
- Написать код, связывающий GRPC сервер с методами доменной области.
- Критерии оценки: Все методы должны быть реализованы
- Бизнес логика (пакет internal/domain в примере) НЕ должен зависеть от кода GRPC сервера
- GRPC-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

Код должен проходить проверки go vet и golint

У преподавателя должна быть возможность заново сгенерировать код по Protobuf спекам

У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install