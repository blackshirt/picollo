package picollo

import (
	"context"
	"picollo/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) OpdItem() OpdItemResolver {
	return &opdItemResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) RencanaWaktu() RencanaWaktuResolver {
	return &rencanaWaktuResolver{r}
}

type opdItemResolver struct{ *Resolver }

func (r *opdItemResolver) Rups(ctx context.Context, obj *model.OpdItem) ([]*model.RupItem, error) {
	panic("not implemented")
}
func (r *opdItemResolver) Tahun(ctx context.Context, obj *model.OpdItem) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ViewOpd(ctx context.Context, id string) (*model.OpdItem, error) {
	panic("not implemented")
}
func (r *queryResolver) ViewRup(ctx context.Context, options *model.RupOptions) ([]*model.RupItem, error) {
	panic("not implemented")
}

type rencanaWaktuResolver struct{ *Resolver }

func (r *rencanaWaktuResolver) Pemilihan(ctx context.Context, obj *model.RencanaWaktu) (*Planning, error) {
	panic("not implemented")
}
func (r *rencanaWaktuResolver) Pelaksanaan(ctx context.Context, obj *model.RencanaWaktu) (*Planning, error) {
	panic("not implemented")
}
func (r *rencanaWaktuResolver) Pemanfaatan(ctx context.Context, obj *model.RencanaWaktu) (*Planning, error) {
	panic("not implemented")
}
