package region

import (
	"odin_tool_v3/libs/context"
	"odin_tool_v3/models/api/master"
)

func WorldList(c *context.Context) {
	worlds := master.GetAllWorlds()

	c.Data["worlds"] = worlds
	c.HTML(200, "region/worldlist")
}
