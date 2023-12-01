package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tlkm-api/configs"
	"tlkm-api/driver/postgre"
	rProductPG "tlkm-api/internal/repository/product/postgre"
	rProductReviewPG "tlkm-api/internal/repository/productreview/postgre"
	sProduct "tlkm-api/internal/service/product"
	sProductV0 "tlkm-api/internal/service/product/v0"
	sProductReview "tlkm-api/internal/service/productreview"
	sProductReviewV0 "tlkm-api/internal/service/productreview/v0"
)

type Server struct {
	client        *http.Server
	preMiddleware []alice.Constructor
	router        *httprouter.Router
}

type InitAttribute struct {
	ServiceProduct       sProduct.Service
	ServiceProductReview sProductReview.Service
}

func InitHttp(config *configs.Config) InitAttribute {
	dbTelkomRead := postgre.GetDBTelkomRead()

	dbTelkomWrite := postgre.GetDBTelkomWrite()

	repoProductPostgre := rProductPG.New(rProductPG.InitAttribute{
		DB: rProductPG.DBList{
			TelkomRead:  dbTelkomRead,
			TelkomWrite: dbTelkomWrite,
		},
	})

	repoProductReviewPostgre := rProductReviewPG.New(rProductReviewPG.InitAttribute{
		DB: rProductReviewPG.DBList{
			TelkomRead:  dbTelkomRead,
			TelkomWrite: dbTelkomWrite,
		},
	})

	serviceProductV0 := sProductV0.New(sProductV0.InitAttribute{
		Repo: sProductV0.RepoAttribute{
			ProductPostgre:       repoProductPostgre,
			ProductReviewPostgre: repoProductReviewPostgre,
		},
	})

	serviceProductReviewV0 := sProductReviewV0.New(sProductReviewV0.InitAttribute{
		Repo: sProductReviewV0.RepoAttribute{
			ProductReviewPostgre: repoProductReviewPostgre,
		},
	})

	httpServer := InitAttribute{
		ServiceProduct:       serviceProductV0,
		ServiceProductReview: serviceProductReviewV0,
	}

	return httpServer
}

func NewHTTP() *Server {
	router := httprouter.New()
	s := &Server{
		router: router,
		client: &http.Server{
			ReadTimeout:  time.Duration(configs.Get().TimeoutHTTP.Read) * time.Second,
			WriteTimeout: time.Duration(configs.Get().TimeoutHTTP.Write) * time.Second,
			Handler:      router,
		},
	}

	s.checker()
	return s
}

func (s *Server) Run() error {
	idleConnClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Printf("HTTP server Shutdown: %v\n", err)
		}
		close(idleConnClosed)
	}()

	fmt.Printf("HTTP server running on port %s\n", os.Getenv("PORT"))

	if err := s.Serve(); err != http.ErrServerClosed {
		return err
	}

	<-idleConnClosed
	fmt.Println("HTTP server shutdown gracefully")
	return nil
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		return err
	}

	return s.client.Serve(lis)
}
func (s *Server) checker() {
	s.AddRoute(http.MethodGet, "/application/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte("OK"))
	})
}

func (s *Server) AddRoute(method, path string, handler http.HandlerFunc) {
	s.router.Handler(method, path, alice.New(s.preMiddleware...).Then(handler))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.client.Shutdown(ctx)
}
