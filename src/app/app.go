/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/7/29 1:28
 * @version     v1.0
 * @filename    app.go
 * @description
 ***************************************************************************/
package app

import (
	"cms/src/config"
	"cms/src/core"
	"cms/src/core/csrf"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	StaticDir  string
	ConfigDir  string
	ConfigFile string
}

type Option func(*options)

func SetStaticDir(staticDir string) Option {
	return func(opts *options) {
		opts.StaticDir = staticDir
	}
}

func SetConfigDir(configDir string) Option {
	return func(opts *options) {
		opts.ConfigDir = configDir
	}
}

func SetConfigFile(configFile string) Option {
	return func(opts *options) {
		opts.ConfigFile = configFile
	}
}

func InitConfig(options *options) error {
	cfg := config.GetInstanceOfConfig()
	err := cfg.Load(options.ConfigDir, options.ConfigFile)
	ctxProj := core.GetInstanceOfContext()
	ctxProj.RunMode = cfg.RunMode
	ctxProj.Http.Host = cfg.Http.Host
	ctxProj.Http.Protocol = cfg.Http.Protocol
	ctxProj.FEAdmin.LocalProtocol = cfg.FEAdmin.FEAdminHttp.LocalProtocol
	ctxProj.FEAdmin.RemoteProtocol = cfg.FEAdmin.FEAdminHttp.RemoteProtocol
	ctxProj.FEAdmin.Domain = cfg.FEAdmin.FEAdminHttp.Domain
	ctxProj.FEAdmin.Host = cfg.FEAdmin.FEAdminHttp.Host
	ctxProj.FEAdmin.Port = cfg.FEAdmin.FEAdminHttp.Port
	ctxProj.FEMain.LocalProtocol = cfg.FEMain.FEMainHttp.LocalProtocol
	ctxProj.FEMain.RemoteProtocol = cfg.FEMain.FEMainHttp.RemoteProtocol
	ctxProj.FEMain.Domain = cfg.FEMain.FEMainHttp.Domain
	ctxProj.FEMain.Host = cfg.FEMain.FEMainHttp.Host
	ctxProj.FEMain.Port = cfg.FEMain.FEMainHttp.Port
	ctxProj.JWTSecret = "liuzhao"
	prk, puk, _ := core.GenRsaKeyPair(2048)
	ctxProj.PublicKey = puk
	ctxProj.PrivateKey = prk
	publicKeyStr, _ := core.PublicKeyToString()
	ctxProj.PublicKeyStr = publicKeyStr
	sessionManager, _ := csrf.NewSessionManager("sProvider", int64(time.Hour*24*7))
	core.GetInstanceOfContext().SessionManager = sessionManager
	go core.GetInstanceOfContext().SessionManager.SessionGC()
	if err != nil {
		logger.WithFields(logger.Fields{
			"directory":  options.ConfigDir,
			"file_names": options.ConfigFile,
		}).Fatal("Configuration files missing. ", err)
	}
	return err
}

func InitServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.GetInstanceOfConfig()
	//addr := fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	addr := fmt.Sprintf(":%s", cfg.Http.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		logger.WithContext(ctx).Printf("HTTP server is running at %s.", addr)
		//server.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
		//err = server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server error. ", err)
		}
	}()
	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.Http.ShutdownTimeout))
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}
}

func InitDB() (*gorm.DB, func(), error) {
	fmt.Println("DB initialisation is ready.")
	cfg := config.GetInstanceOfConfig()
	db, clean, err := cfg.NewDB()
	if err != nil {
		logger.Fatal("DB initialisation failed. ", err)
		return nil, clean, err
	}
	err = cfg.AutoMigrate(db)
	if err != nil {
		logger.Fatal("DB migration failed. ", err)
		return nil, clean, err
	}
	return db, clean, err
}

func Init(ctx context.Context, opts ...Option) func() {
	cfg := config.GetInstanceOfConfig()
	// initialising options
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}
	// init config
	err := InitConfig(&options)
	if err != nil {
		logger.Fatal("Init config failed. ", err)
	}
	// init injector
	injector, injectorClean, _ := InitInjector()
	//injector.DB = db
	logger.WithFields(logger.Fields{
		"db_type":   cfg.DB.Type,
		"db_name":   cfg.Mysql.DBName,
		"user_name": cfg.Mysql.UserName,
		"host":      cfg.Mysql.Host,
		"port":      cfg.Mysql.Port,
	}).Info("DB connected.")
	// init server
	serverClean := InitServer(ctx, injector.Engine)
	return func() {
		serverClean()
		injectorClean()
	}
}

func Launch(ctx context.Context, opts ...Option) {
	cfg := config.GetInstanceOfConfig()
	clean := Init(ctx, opts...)
	logger.WithFields(logger.Fields{
		"app_name":         cfg.App.AppName,
		"version":          cfg.App.Version,
		"pid":              os.Getpid(),
		"protocol":         cfg.Http.Protocol,
		"protocol_version": cfg.Http.ProtocolVersion,
		"host":             cfg.Http.Host,
		"port":             cfg.Http.Port,
	}).Info("Service launched.")
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
LOOP:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("Interrupt received [%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break LOOP
		case syscall.SIGHUP:
		default:
			break LOOP
		}
	}
	defer logger.WithContext(ctx).Infof("Server is shutting down.")
	defer time.Sleep(time.Second)
	defer os.Exit(state)
	defer clean()
}
