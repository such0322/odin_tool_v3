package misc

import (
	"time"
)

const GiftCodeStep = 100

type GiftCode struct {
	Model
	Code        string
	Batch       int
	Channel     string
	Type        int
	Quantity    int
	Package     string
	Status      int
	StartDate   time.Time
	EndDate     time.Time
	LastModDate time.Time
}

func GetGiftCodeByPage(pager int) []GiftCode {
	var giftcodes []GiftCode
	offset := 0
	if pager > 0 {
		offset = (pager - 1) * GiftCodeStep
	}

	DB.Limit(GiftCodeStep).Offset(offset).Find(&giftcodes)

	return giftcodes
}

func GetGiftCodeCount() (count int) {
	var gc GiftCode
	DB.Model(&gc).Count(&count)
	return count
}
