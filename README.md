Калькулятор (Calc API)

Описание

Этот проект представляет собой веб-сервис для вычисления арифметических выражений, написанный на языке Go.

Пользователь отправляет POST-запрос с арифметическим выражением, и сервис возвращает результат вычисления в формате JSON.

Проект поддерживает базовые операции: сложение, вычитание, умножение, деление и скобки.

Структура проекта

cmd/calc_service/main.go - точка входа в приложение.

internal/application/application.go - основной код сервиса, обработка HTTP-запросов.

pkg/calculation/calculator.go - логика калькулятора.

pkg/calculation/calculator_test.go - тесты для калькулятора.

Запуск проекта

Склонируйте репозиторий:

git clone https://github.com/lollmark/calc_go.git

Перейдите в директорию проекта:

cd calc_go

Запустите сервис:

go run ./cmd/calc_service/...

Сервис будет доступен по адресу http://localhost:8080.

Пример использования

POST /api/v1/calculate

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Ответ:

{
  "result": 6
}

Обработка ошибок

Невалидное выражение:

{
  "error": "Expression is not valid"
}

Внутренняя ошибка сервера:

{
  "error": "Internal server error"
}

Тестирование

Запуск тестов:

go test ./...

Зависимости

Go версии 1.20 и выше.

Лицензия

MIT License.

