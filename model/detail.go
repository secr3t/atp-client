package model

import "strings"

type DetailResult struct {
	Item  *DetailItem `json:"item"`
	Error string      `json:"error"`
}

func (dr DetailResult) IsError() bool {
	return dr.Error != ""
}

type ItemImgs struct {
	Url string `json:"url"`
}

type Sku struct {
	Price          string `json:"price"`
	TotalPrice     int    `json:"total_price"`
	OrginalPrice   string `json:"orginal_price"`
	Properties     string `json:"properties"`
	PropertiesName string `json:"properties_name"`
	Quantity       string `json:"quantity"`
	SkuID          string `json:"sku_id"`
}

type Skus struct {
	Sku []Sku `json:"sku"`
}

type DetailItem struct {
	NumIid    string            `json:"num_iid"`
	Title     string            `json:"title"`
	Price     string            `json:"price"`
	DetailURL string            `json:"detail_url"`
	PicURL    string            `json:"pic_url"`
	Desc      string            `json:"desc"`
	ItemImgs  []ItemImgs        `json:"item_imgs"`
	Skus      Skus              `json:"skus"`
	PropsList map[string]string `json:"props_list"`
	PropsImg  map[string]string `json:"props_img"`
	DescImg   []string          `json:"desc_img"`
}

type Option struct {

}

func (di DetailItem) GetItemImgs() []string {
	imgs := make([]string, 0)
	for _, itemImgs := range di.ItemImgs {
		img := itemImgs.Url
		imgs = append(imgs, ValidImg(img))
	}

	return imgs
}

func (di DetailItem) GetPropImg(propPath string) string {
	img := di.PropsImg[propPath]

	return ValidImg(img)
}


func ValidImg(img string) string  {
	if img == "" {
		return ""
	}
	if !strings.HasPrefix(img, "http") {
		img = "http:" + img
	}
	return img
}