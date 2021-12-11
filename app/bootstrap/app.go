package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/routes"
	"github.com/gofiber/fiber/v2"

	sentryfiber "github.com/directoryxx/fiber-clean-template/pkg/sentry"
	"github.com/getsentry/sentry-go"
)

const idleTimeout = 5 * time.Second

// Dispatch is handle routing
func Dispatch(ctx context.Context, log interfaces.Logger, enforcer *casbin.Enforcer) {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	fmt.Println(os.Getenv("SENTRY_DSN"))

	errSentry := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1,
		Debug:            false,
		AttachStacktrace: true,
		// BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
		// 	if hint.Context != nil {
		// 		if ctx, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
		// 			// You have access to the original Context if it panicked
		// 		}
		// 	}
		// 	return event
		// },
	})

	sentryHandler := sentryfiber.New(sentryfiber.Options{})

	if errSentry != nil {
		panic(errSentry)
	}

	app.Use(sentryHandler)

	app.Static("/storage/", "/app/public/")

	// app.Use(pprof.New())
	routes.RegisterRoute(app, ctx, log, enforcer)

	go func() {
		if errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT")); errApp != nil {
			log.LogError("%s", errApp)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// sqlHandler.Close()
	// redisHandler.Close()
	fmt.Println("Fiber was successful shutdown.")

	// if errApp != nil {
	// 	log.LogError("%s", errApp)
	// }
}
