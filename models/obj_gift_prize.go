package models

type ObjGiftPrize struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	PrizeNum   int    `json:"_"`
	LeftNum    int    `json:"_"`
	PrizeCodeA int    `json:"_"`
	PrizeCodeB int    `json:"_"`
	Img        string `json:"img"`
	PrizeOrder int    `json:"prize_order"`
	Gtype      int    `json:"gtype"`
	Gdata      string `json:"gdata"`
}
