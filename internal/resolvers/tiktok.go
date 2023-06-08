package resolvers

import (
	"context"
	"strings"

	"github.com/dollarkillerx/graphql_template/internal/generated"
	"github.com/dollarkillerx/graphql_template/internal/utils/tiktok"
)

// // foo
func (r *queryResolver) App(ctx context.Context, appID string) (*generated.AppInfo, error) {
	return &generated.AppInfo{
		AppID:              appID,
		AppVersion:         0,
		MinimumVersion:     0,
		State:              generated.AppStateEnable,
		ErrorNotification:  "",
		NormalNotification: "",
	}, nil
}

// // foo
func (r *queryResolver) VideoDownload(ctx context.Context, url string) (*generated.VideoInfo, error) {
	url = strings.TrimSpace(url)
	img, videos, err := tiktok.TiktokGet(url, false)
	if err != nil {
		img, videos, err = tiktok.TiktokGet(url, true)
		if err != nil {
			return nil, err
		}
	}

	return &generated.VideoInfo{
		Img:  img,
		Urls: videos,
	}, nil
}
