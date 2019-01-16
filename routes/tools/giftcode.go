package tools

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"odin_tool_v3/libs/context"
	"odin_tool_v3/libs/pages"
	"odin_tool_v3/models/api/master/bonus"
	"odin_tool_v3/models/bridge/misc"
	"strconv"
	"time"
)

func DownloadCode(c *context.Context) {
	batch, _ := strconv.Atoi(c.Req.FormValue("batch"))

	gcs := misc.GetGiftCodeByBatch(batch)

	filename := "odinGiftCode" + strconv.Itoa(batch) + ".csv"
	b := &bytes.Buffer{}
	b.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(b)
	w.Write([]string{"礼包码"})
	for _, vo := range gcs {
		w.Write([]string{
			vo.Code,
		})
	}
	w.Flush()

	c.Header().Set("Content-Type", "text/csv")
	c.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	c.RawData(200, b.Bytes())

}

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
