package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/pz3-http/internal/api"
	"example.com/pz3-http/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Коллекция
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)

	// Элемент
	mux.HandleFunc("GET /tasks/", h.GetTask)
	// Маршрут метод PATCH
	mux.HandleFunc("PATCH /tasks/", h.UpdateTask)
	// Маршрут метод DELETE
	mux.HandleFunc("DELETE /tasks/", h.DeleteTask)

	// Подключаем middleware (CORS -> Logging)
	handler := api.CORS(api.Logging(mux))

	// Создаем HTTP сервер с настройками
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Запускаем сервер в отдельной goroutine
	go func() {
		log.Println("listening on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Канал для сигналов ОС
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Ждем сигнал завершения
	<-quit
	log.Println("Shutting down server...")

	// Создаем контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Останавливаем сервер
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
