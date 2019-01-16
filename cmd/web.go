package cmd

import (
	"odin_tool_v3/libs/context"
	"odin_tool_v3/libs/logger"
	"odin_tool_v3/libs/setting"
	"odin_tool_v3/libs/template"
	"odin_tool_v3/models"
	"odin_tool_v3/routes"
	"odin_tool_v3/routes/auth"
	"odin_tool_v3/routes/index"
	"odin_tool_v3/routes/region"
	"odin_tool_v3/routes/tools"

	"github.com/go-macaron/macaron"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
)

var Web = cli.Command{
	Name:   "web",
	Usage:  "后台网页工具",
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "port,p",
			Value: "3000",
			Usage: "链接用的端口",
		},
	},
}

func runWeb(c *cli.Context) error {
	//设置日志启用
	logger.InitLogger()
	//使用经典的macaron实例
	m := macaron.New()
	//m.Use(logger.SetExectimeLog())
	//m.Use(macaron.Logger())
	m.Use(logger.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))

	funcMap := template.NewFuncMap()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		IndentJSON: false,
		Funcs:      funcMap,
	}))
	setting.LoadCfg()
	models.NewEngines()

	sessionOptions := session.Options{
		Provider: "memory",
		//ProviderConfig:"",
		CookieName:  "OdinToolSession",
		Gclifetime:  3600,
		Maxlifetime: 86400,
	}
	m.Use(session.Sessioner(sessionOptions))
	m.Use(context.Contexter())

	router(m)
	port := 4000
	if c.IsSet("port") {
		port = c.Int("port")
	}

	m.Run(port)
	return nil
}

func router(m *macaron.Macaron) {
	//needLogin:=context.NeedLogin()

	//路由
	m.Get("/", index.Index)
	m.Get("/debug", index.Debug)

	m.Get("/auth/login", auth.Login)
	m.Post("/auth/postLogin", auth.PostLogin)
	m.Get("/auth/logout", auth.Logout)

	m.Group("", func() {
		m.Get("/worlds", region.WorldList)
		m.Get("/giftcodes", tools.GiftCodeList)
		m.Get("/gift/new", tools.NewGift)
		m.Post("/gift/new", tools.CreateGift)
		m.Get("/gift/randomCode", tools.RandomCode)
		m.Get("/gift/getBounsAll", tools.GetBounsAll)
		m.Get("/gift/download", tools.DownloadCode)
	})

	m.NotFound(routes.NotFound)
}
