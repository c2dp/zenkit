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
	picpath := "data/pic/"
	picname := guid.S() + ".jpg"
	res, err := g.Client().Get(ctx, image.Url)
	if err != nil {
		return err
	}
	defer res.Close()
	gfile.PutBytes(picpath+picname, res.ReadAll())
	g.Log().Debugf(ctx, "picname: %s", picname)

	return nil

}
