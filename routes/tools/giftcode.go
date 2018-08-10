package tools

import (
	"odin_tool_v3/libs/context"
	"odin_tool_v3/libs/pages"
	"odin_tool_v3/models/bridge/misc"
	"strconv"
)

func GiftCodeList(c *context.Context) {
	pager, err := strconv.Atoi(c.Req.FormValue("pager"))
	if err != nil {
		pager = 1
	}
	giftcodes := misc.GetGiftCodeByPage(pager)

	count := misc.GetGiftCodeCount()
	pages := &pages.Pages{Count: count, Page: pager, PrePage: misc.GiftCodeStep, Url: "giftcodes"}

	c.Data["CodeList"] = giftcodes
	c.Data["Pages"] = pages.Get()
	c.HTML(200, "tools/giftcode/list")

}
