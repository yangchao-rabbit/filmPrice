package apps

import (
	"filmPrice/config"
	_ "filmPrice/docs"
	systemDao "filmPrice/internal/apps/system/dao"
	"filmPrice/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"regexp"
	"strings"
)

var (
	// Log : apps 全局日志
	Log *log.Logger
	// ginApps gin应用存储map
	ginApps = map[string]GinService{}

	implApps = map[string]ImplService{}
)

type GinService interface {
	Name() string
	Registry(gin.IRouter)
}

type ImplService interface {
	Name() string
	Init() error
}

// NewGinServer 创建Gin服务
func NewGinServer() (*gin.Engine, error) {
	Log.Printf("\n%v GinServer启动...\n", config.Get().App.URL())

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.Logger(Log))
	r.Use(gin.Recovery())

	// 设置api路由
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	setApiRouter(api)

	autoRegisterRoutesToPerms(r.Routes())

	// swagger
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "VP_SWAGGER"))

	// 展示路由信息
	showRouting(r.Routes())

	return r, nil
}

// RegistryGin 注册应用到 ginApps
func RegistryGin(svc GinService) {
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("GinService: %v 已注册", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

func RegistryImpl(i ImplService) {
	if _, ok := implApps[i.Name()]; ok {
		panic(fmt.Sprintf("ImplService: %v 已注册", i.Name()))
	}

	implApps[i.Name()] = i
}

// GetImplSvc 获取实现svc
func GetImplSvc(name string) any {
	if v, ok := implApps[name]; ok {
		return v
	}

	panic(name + "not found")
}

func InitImplApps(logger *log.Logger) error {
	Log = logger

	for _, v := range implApps {
		if err := v.Init(); err != nil {
			fmt.Println(err)
			return err
		}
		Log.Printf("初始化 %v ImplService成功", v.Name())
	}
	return nil
}

func setApiRouter(r gin.IRouter) {
	// 循环注册所有HTTP路由
	for _, v := range ginApps {
		v.Registry(r.Group(v.Name()))

		Log.Printf("注册 %v 应用路由成功", v.Name())
	}
}

// showRouting 打印路由信息
func showRouting(routes gin.RoutesInfo) {
	Log.Println("=============路由信息============")
	for _, v := range routes {
		Log.Printf("Method: %v | Path: %v | Handler: %v", v.Method, v.Path, v.Handler)
	}
	Log.Println("=============路由信息============")
}

// 自动注册权限
func autoRegisterRoutesToPerms(routes gin.RoutesInfo) {
	for _, v := range routes {
		// 忽略OPTIONS请求
		if v.Method == "OPTIONS" {
			continue
		}

		// 忽略以/swagger开头的请求
		if strings.HasPrefix(v.Path, "/swagger") {
			continue
		}

		// 忽略以/static开头的请求
		if strings.HasPrefix(v.Path, "/static") {
			continue
		}

		perm := systemDao.SystemPermModel{
			Name:   ExtractFuncName(v.Handler),
			Path:   v.Path,
			Method: v.Method,
		}

		var count int64
		config.GetDB().Model(&systemDao.SystemPermModel{}).
			Where("method = ? AND path = ?", perm.Method, perm.Path).
			Count(&count)

		if count == 0 {
			config.GetDB().Create(&perm)
		} else {
			config.GetDB().Model(&perm).
				Where("method = ? AND path = ?", perm.Method, perm.Path).
				Updates(&perm)
		}
	}
}

// ExtractFuncName 提取 Handler 最后一个函数名称
func ExtractFuncName(handler string) string {
	// 正则匹配以 funcX 或方法名结尾的结构
	re := regexp.MustCompile(`\.(\w+)(\.func\d+)?$`)
	matches := re.FindStringSubmatch(handler)
	if len(matches) > 1 {
		return matches[1] // 提取 Delete
	}
	return handler
}
