package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv библиотека

	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enable")

	// TODO: init storage инициализируем сторадже: sqllite

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed not init storage", sl.Err(err))
		os.Exit(1)
	}
	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
	}

	log.Info("saved url", slog.Int64("id", id))

	id, err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
	}

	_ = storage
	//добавление урла в БД
	//вывод инфы о id
	//ошибка при добавении одикаковых alias

	// TODO: init router: chi, render

	// TODO: run server
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
