package fs

import (
	"context"
	"strings"

	"github.com/NodeSeekDev/nlist/v3/internal/model"
	"github.com/NodeSeekDev/nlist/v3/internal/op"
	"github.com/NodeSeekDev/nlist/v3/server/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func link(ctx context.Context, path string, args model.LinkArgs) (*model.Link, model.Obj, error) {
	storage, actualPath, err := op.GetStorageAndActualPath(path)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "failed get storage")
	}
	l, obj, err := op.Link(ctx, storage, actualPath, args)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "failed link")
	}
	if l.URL != "" && !strings.HasPrefix(l.URL, "http://") && !strings.HasPrefix(l.URL, "https://") {
		if c, ok := ctx.(*gin.Context); ok {
			l.URL = common.GetApiUrl(c.Request) + l.URL
		}
	}
	return l, obj, nil
}
