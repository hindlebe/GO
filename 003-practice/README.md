## Практическое занятие №3 Реализация простого HTTP-сервера на стандартной библиотеке net/http. Обработка запросов GET/POST

### ФИО: Мальцев Никита Михайлович
### Группа: ЭФМО 02-25

### 1. Структура проекта

```
003-practice/
│   README.md
│   temp.json
│
├───photos                          #Дериктория для фото
└───pz3-http                        #Основной проект
    │   go.mod                      #Модуль Go
    │
    ├───cmd
    │           main.go             #Точка для входа приложения
    │
    └───internal
        ├───api
        │       handlers.go         # Обработчики HTTP запросов
        │       handlers_test.go    # Юнит-тесты
        │       middleware.go       # Middleware (CORS, логирование)
        │       responses.go        # Вспомогательные функции ответов
        │
        └───storage
                memory.go            # In-memory хранилище задач
```

1. Health check

![Health check](./photos/Heath_check.png)

2. Создание и просмотр задач

![tasks](./photos/tasks.png)

3. Получение задачи по `ID`

![ID TASK](./photos/ID.png)

### Функциональность

- `GET /health` - проверка работоспособности сервера

- `GET /tasks` - список всех задач (с фильтрацией ?q=)

- `POST /tasks` - создание новой задачи

- `GET /tasks/{id}` - получение задачи по ID

- `PATCH /tasks/{id}` - обновление задачи

- `DELETE /tasks/{id}` - удаление задачи

### Дополнительно

1. `CORS (минимально)`: добавить заголовки Access-Control-Allow-Origin: * для GET/POST (в отдельной middleware).

![CORS (минимально)](./photos/п1.png)

2. Валидация длины `title` (например, 1…140 символов).

![Валидация длины title (например, 1…140 символов).](./photos/п2.png)

3. Метод `PATCH /tasks/{id}`для отметки Done=true.

![Метод PATCH /tasks/{id} для отметки Done=true.](./photos/п3.png)

4. Метод `DELETE /tasks/{id}`

    - Задача для удаления \ проверка

    ![TASK CREATE AND CHECK](./photos/п4_1.png)

    - Удаление задачи \ проверка

    ![DELETE THE TASK AND CHECK](./photos/п4_2.png)

5. `Graceful shutdown` через http.Server и контекст.

![Graceful shutdown](./photos/п5.png)

6. `Юнит-тесты` обработчиков с `httptest`

![Юнит-тесты](./photos/п6.png)