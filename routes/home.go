package routes

import "odin_tool_v3/libs/context"

func NotFound(c *context.Context) {
	c.Data["Title"] = "Page Not Found"
	c.NotFound()
}
