package services

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type UploadConfigResp struct {
	UpToken string `json:"uptoken"`
}

func (s *Service) UploadConfig(ctx *gin.Context) (resp UploadConfigResp, err error) {
	accessKey := "irrxGB4boUnvhiwKa8bB0DWSyjRhZVTS4mCPJ4bs"
	secretKey := "drujaYApyeUzJA3V2QTXP4nW_gvNtEO-rxpfBwZJ"

	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "haolinjiu"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)

	resp = UploadConfigResp{
		UpToken: putPolicy.UploadToken(mac),
	}

	return
}
