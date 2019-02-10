package main

import (
	"context"
	"os"
	"os/signal"
	"time"
	"todo/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := route.Init()

	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	db, err := gorm.Open("mysql", "root:@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down server")
		}
	}()

	// Wait for signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
