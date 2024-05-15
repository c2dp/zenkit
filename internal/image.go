package internal

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/guid"
)

type ImageItem struct {
	Url string `json:"PicUrl"`
}

func DownloadImage(ctx context.Context, image ImageItem) error {
	res, err := g.Client().Get(ctx, image.Url)
	if err != nil {
		return err
	}
	defer res.Close()
	gfile.PutBytes("data/pic/"+guid.S()+".jpg", res.ReadAll())

	return nil

}
