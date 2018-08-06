package cmd

import (
	"log"
	"odin_tool_v3/libs/context"
	"odin_tool_v3/libs/setting"
	"odin_tool_v3/models"
	"odin_tool_v3/routes"
	"odin_tool_v3/routes/auth"
	"odin_tool_v3/routes/index"
	"os"

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

	//m := macaron.New()
	//m.Use(macaron.Logger())
	//m.Use(macaron.Recovery())
	//m.Use(macaron.Static("./public", macaron.StaticOptions{SkipLogging: true}))

	//使用经典的macaron实例
	m := macaron.New()
	m.Use(globalInit())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer(macaron.RenderOptions{IndentJSON: false}))
	setting.LoadCfg()
	models.NewEngine()

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

	m.Run()
	return nil
}

func router(m *macaron.Macaron) {
	//路由
	m.Get("/", index.Index)
	m.Get("debug", index.Debug)

	m.Get("auth/login", auth.Login)
	m.Post("auth/postLogin", auth.PostLogin)

	m.NotFound(routes.NotFound)
}

func globalInit() macaron.Handler {
	return func(c *macaron.Context) {
		//fixme 或许就不应该在这里指定logger的位置,应该使用其他的包比如clog？
		logPath := "logs/app.log"

		logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("open logfile failed")
		}
		c.Map(logfile)
		//todo 这里close了下面都写不进去了，暂时还不知道应该在哪里关闭
		//defer logfile.Close()
		logger := log.New(logfile, "[DEBUG]", log.LstdFlags|log.Llongfile)
		c.Map(logger)

	}
}
