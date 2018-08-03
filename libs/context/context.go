package context

import (
	"fmt"
	"net/http"
	"odin_tool_v3/models"

	"github.com/go-macaron/macaron"
	"github.com/go-macaron/session"
)

type Context struct {
	*macaron.Context
	Session session.Store

	User *models.User
}

//todo 这边的依赖注入还不是很懂，暂时这样可以使用,有点懂了
func NewContexter() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		c := &Context{
			Context: ctx,
			Session: sess,
		}

		c.User = models.UserSignin(sess)
		ctx.Map(c)
	}
}

func (c *Context) NotFound() {
	c.Handle(http.StatusNotFound, "", nil)
}

func (c *Context) ServerError(title string, err error) {
	c.Handle(http.StatusInternalServerError, title, err)
}

func (c *Context) Handle(status int, title string, err error) {
	switch status {
	case http.StatusNotFound:
		c.Data["Title"] = "Page Not Found"
	case http.StatusInternalServerError:
		c.Data["Title"] = "Internal Server Error"
		//log.Error(3, "%s: %v", title, err)
		//if !setting.ProdMode || (c.IsLogged && c.User.IsAdmin) {
		//	c.Data["ErrorMsg"] = err
		//}
	}
	c.HTML(status, fmt.Sprintf("status/%d", status))
}
