package picollo

import (
	"context"
	"picollo/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Service model.Service
}

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
	rups := make([]*model.RupItem, 0)
	res, err := r.Service.Rup(ctx, *options)

	// loop SALAH: ini cuma akan menghasilkan rups berisi address yang sama
	// for _, item := range res {
	// rups = append(rups, &item)
	// }
	for i, _ := range res {
		rups = append(rups, &res[i])
	}

	return rups, err
}

type rupRekapItemResolver struct{ *Resolver }

func (r *rupRekapItemResolver) Rups(ctx context.Context, obj *model.RupRekapItem) ([]*model.RupItem, error) {
	panic("not implemented")
}
