package picollo

import (
	"context"
	"picollo/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) RupRekapItem() RupRekapItemResolver {
	return &rupRekapItemResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ViewOpd(ctx context.Context, id string) (*model.RupRekapItem, error) {
	panic("not implemented")
}
func (r *queryResolver) ViewRup(ctx context.Context, options *model.RupOptions) ([]*model.RupItem, error) {
	panic("not implemented")
}

type rupRekapItemResolver struct{ *Resolver }

func (r *rupRekapItemResolver) Rups(ctx context.Context, obj *model.RupRekapItem) ([]*model.RupItem, error) {
	panic("not implemented")
}
func (r *rupRekapItemResolver) Tahun(ctx context.Context, obj *model.RupRekapItem) (string, error) {
	panic("not implemented")
}
