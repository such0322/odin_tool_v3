package tools

import (
	"encoding/json"
	"net/http"
	"odin_tool_v3/libs/context"
	"odin_tool_v3/libs/pages"
	"odin_tool_v3/models/api/master/bonus"
	"odin_tool_v3/models/bridge/misc"
	"strconv"
	"time"
)

func GiftCodeList(c *context.Context) {
	pager, err := strconv.Atoi(c.Req.FormValue("pager"))
	if err != nil {
		pager = 1
	}
	giftcodes := misc.GetGiftCodeByPage(pager)

	count := misc.GetGiftCodeCount()
	pages := &pages.Pages{Count: count, Page: pager, PrePage: misc.GIFTCODE_STEP, Url: "giftcodes"}

	c.Data["CodeList"] = giftcodes
	c.Data["Pages"] = pages.Get()
	c.HTML(200, "tools/giftcode/list")

}

func NewGift(c *context.Context) {
	c.Data["RewardType"] = bonus.RewardType
	c.HTML(200, "tools/giftcode/new")
}

func CreateGift(c *context.Context) {
	r := c.Req
	r.ParseForm()

	data := make(map[string]interface{})
	data["code"] = r.FormValue("code")
	data["type"], _ = strconv.Atoi(r.FormValue("type"))
	data["channel"] = r.FormValue("channel")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	if quantity <= 0 {
		//TODO 做个flash，跳转后弹出错误信息
		c.Redirect("/gift/new", http.StatusFound)
		return
	}
	data["quantity"] = quantity
	reward_types := r.Form["reward_type"]
	reward_ids := r.Form["reward_id"]
	reward_qs := r.Form["reward_quantity"]
	bd := bonus.BonusData{}
	bds := bd.CreateRewards(reward_types, reward_ids, reward_qs)
	t := time.Now()
	data["batch"] = t.Unix()
	if len(bds) == 0 {
		data["package"] = "[]"
	} else {
		data["package"], _ = json.Marshal(bds)
	}
	data["start_date"] = r.FormValue("start_date")
	data["end_date"] = r.FormValue("end_date")
	misc.CreateGiftCodes(data)

	c.Redirect("/giftcodes", http.StatusFound)
}

func RandomCode(c *context.Context) {
	code := misc.GetRandomCode()
	c.PlainText(200, []byte(code))
}

func GetBounsAll(c *context.Context) {
	rewardType := c.Req.FormValue("reward_type")
	bonus, err := bonus.NewBonus(rewardType)
	if err != nil {
		c.WriteHeader(404)
		return
	}
	bonus.GetRewardNames()
	c.JSON(200, bonus)
}
