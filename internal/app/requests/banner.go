package requests

type BannerListReq struct {
	VillageId int16 `json:"village_id" form:"village_id" validate:"required"`
}

type BannerListResp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ImgUrl    string `json:"img_url"`
	JumpUrl   string `json:"jump_url"`
	Sort      int16  `json:"sort"`
	VillageId int16  `json:"village_id"`
}
