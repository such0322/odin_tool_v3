package index

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"odin_tool_v3/libs/context"
	"odin_tool_v3/models/tool"
	"strconv"
)

func Index(c *context.Context, log *log.Logger) {
	if c.User != nil {
		c.Data["Name"] = c.User.Name
	} else {
		c.Data["Name"] = "Guest"
	}
	c.HTML(200, "index/index")
	//c.JSON(200, []string{"a", "b"})
}

func Debug(c *context.Context) {
	id, _ := strconv.Atoi(c.Req.FormValue("id"))
	pwd := c.Req.FormValue("pwd")
	md5ctx := md5.New()
	md5ctx.Write([]byte(pwd))
	md5pwd := md5ctx.Sum(nil)
	user := tool.GetUserById(id)

	c.JSON(200, &Aaa{user, hex.EncodeToString(md5pwd)})
}

type Aaa struct {
	*tool.User
	S string
}
