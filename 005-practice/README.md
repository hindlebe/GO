## Практическое занятие №5 Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)

### 1. ФИО: Мальцев Никита Михайлович  
### Группа: ЭФМО-02-25

### 2. Окружение

- Go: go version go1.25.3 windows/amd64

- PostgreSQL: psql (PostgreSQL) 17.6

- ОС: Windows 11 (с использованием WSL2 - MINGW64_NT-10.0-19045)

### 3 Скриншоты

- создание БД/таблицы в psql

![created_db](./photos/created_db.png)

- успешный вывод go run . (вставка и список задач);

![output_go_run](./photos/output1.png)

![output_go_run](./photos/output2.png)

- SELECT * FROM tasks; в psql

![SELECT-QUERY](./photos/SELECT_query.png)

### 4 Код:

- [db.go](./db.go)
- [repository.go](./repository.go)
- [main.go](./main.go)

```GO
//Фрагмент main.go
func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:353722Zz@localhost:5432/todo?sslmode=disable"
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	// Вставка задач
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
	for _, title := range titles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Fatalf("CreateTask error: %v", err)
		}
		log.Printf("Inserted task id=%d (%s)", id, title)
	}

	// Вывод всех задач
	tasks, err := repo.ListTasks(ctx)
	// ... вывод результатов

	// ТЕСТИРОВАНИЕ ПРОВЕРОЧНЫХ ЗАДАНИЙ
	// ... код тестирования ListDone, FindByID, CreateMany
}
```

### 5	Краткие ответы:
- Что такое пул соединений *sql.DB и зачем его настраивать?

**Пул соединений `*sql.DB`** — это набор заранее установленных подключений к БД, которые переиспользуются между запросами. Настройка нужна для оптимизации производительности: ограничение максимального числа соединений предотвращает перегрузку БД, а поддержание соединений в простое уменьшает задержки при новых запросах.

- Почему используем плейсхолдеры $1, $2?

**Плейсхолдеры $1, $2** используются для защиты от SQL-инъекций. Они гарантируют, что пользовательские данные передаются как параметры, а не как часть SQL-кода, что предотвращает выполнение злонамеренного SQL-кода.*


- Чем `Query`, `QueryRow` и `Exec` отличаются?

`Query` — для запросов, возвращающих множество строк (SELECT без WHERE или с условием, возвращающим много записей)

`QueryRow` — для запросов, возвращающих не более одной строки (SELECT с первичным ключом)

`Exec` — для запросов, не возвращающих строки (INSERT, UPDATE, DELETE)

### 6 Обоснование транзакций и настроек пула:

- `SetMaxOpenConns(10)`:

    - Для чего: Ограничение максимального числа одновременных соединений
    - Обоснование: Предотвращение перегрузки БД при высокой нагрузке

- `SetMaxIdleConns(5)`:

    - Для чего: Поддержание соединений в "простое" для быстрого переиспользования
    - Обоснование: Уменьшение задержек на установку нового соединения

- `SetConnMaxLifetime(30 * time.Minute)`:

    - Для чего: Ограничение времени жизни соединения
    - Обоснование: Предотвращение использования "устаревших" соединений

