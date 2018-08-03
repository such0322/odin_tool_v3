package index

import (
	"log"
	"odin_tool_v3/libs/context"
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
