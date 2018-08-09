package region

import (
	"odin_tool_v3/libs/context"
)

func WorldList(c *context.Context) {

	c.HTML(200, "region/worldlist")
}
