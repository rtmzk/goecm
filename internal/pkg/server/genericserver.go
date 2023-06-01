package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go-ecm/internal/goecmagent/static"
	"go-ecm/internal/goecmserver/util/docker"
	"go-ecm/internal/pkg/middleware"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type GenericServer struct {
	middlewares         []string
	mode                string
	SecureServingInfo   *SecureServingInfo
	InSecureServingInfo *InSecureServingInfo
	ShutdownTimeout     time.Duration

	*gin.Engine
	healthz         bool
	enableMetric    bool
	enableProfiling bool

	insecureServer, secureServer *http.Server
}

func initGenericServer(s *GenericServer) {
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

func (s *GenericServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericServer) InstallMiddlewares() {
	s.Use(middleware.RequestID())
	for _, m := range s.middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware: %s", m)
			continue
		}
		log.Infof("install middleware: %s", m)
		s.Use(mw)
	}
	s.Use(middleware.Context())
}

func (s *GenericServer) InstallAPIs() {
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			core.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})
	}

	if s.enableMetric {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	if s.enableProfiling {
		pprof.Register(s.Engine)
	}
}

func (s *GenericServer) Run() error {
	s.insecureServer = &http.Server{
		Addr:    s.InSecureServingInfo.Address,
		Handler: s,
	}

	s.secureServer = &http.Server{
		Addr:    s.SecureServingInfo.Address(),
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("Start to listening the incoming request on http address: %s", s.InSecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.InSecureServingInfo.Address)

		return nil
	})

	eg.Go(func() error {
		key, cert := s.SecureServingInfo.CertKey.KeyFile, s.SecureServingInfo.CertKey.CertFile
		if cert == "" || key == "" || s.SecureServingInfo.BindPort == 0 {
			return nil
		}

		log.Infof("Start to listening the incoming request on https address: %s", s.SecureServingInfo.Address())

		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.SecureServingInfo.Address())

		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if s.healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	// 先安装docker
	if !utils.IsExist("/usr/bin/docker") {
		if utils.IsExist("/tmp/docker.rpm") {
			_ = os.Remove("/tmp/docker.rpm")
		}
		_ = os.WriteFile("/tmp/docker.rpm", static.DOCKERPKG, 0644)
		err := exec.Command("rpm", "-Uvh", "/tmp/docker.rpm").Run()
		if err != nil {
			log.Error("docker install failed.")
			return err
		}
		_ = exec.Command("systemctl", "enable", "docker", "--now").Run()
		log.Info("docker was installed successfully.")

		cli, err := docker.GetDockerClientOr()
		if err != nil {
			return err
		}
		docker.SetClient(cli)
	}

	// 插入swarm join token
	//var tokens v1.SwarmJoin
	//swarmSpec, err := docker.Client().SwarmInspect(context.Background())
	//if err != nil {
	//	return err
	//}

	//tokens.WrkJoinToken = swarmSpec.JoinTokens.Worker
	//tokens.MgrJoinToken = swarmSpec.JoinTokens.Manager
	//err = GetTokens()
	//if err != nil {
	//	errTkn := store.Client().Swarm().GenderToken(context.Background(), &tokens, metav1.CreateOptions{})
	//	if errTkn != nil {
	//		return errTkn
	//	}
	//}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
func (s *GenericServer) RunAgent() error {
	s.insecureServer = &http.Server{
		Addr:    s.InSecureServingInfo.Address,
		Handler: s,
	}

	s.secureServer = &http.Server{
		Addr:    s.SecureServingInfo.Address(),
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("Start to listening the incoming request on http address: %s", s.InSecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.InSecureServingInfo.Address)

		return nil
	})

	eg.Go(func() error {
		key, cert := s.SecureServingInfo.CertKey.KeyFile, s.SecureServingInfo.CertKey.CertFile
		if cert == "" || key == "" || s.SecureServingInfo.BindPort == 0 {
			return nil
		}

		log.Infof("Start to listening the incoming request on https address: %s", s.SecureServingInfo.Address())

		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.SecureServingInfo.Address())

		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if s.healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *GenericServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.secureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown secure server failed: %s", err.Error())
	}

	if err := s.insecureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown insecure server failed: %s", err.Error())
	}
}

func (s *GenericServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.InSecureServingInfo.Address)
	if strings.Contains(s.InSecureServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InSecureServingInfo.Address, ":")[1])
	}

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("The router has been deployed successfully.")

			resp.Body.Close()
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			log.Fatal("can not ping http server within the specified time interval.")
		default:
		}
	}
}
