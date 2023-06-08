package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Resolver struct{}

// // foo
func (r *mutationResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UserRegistration(ctx context.Context, input *UserRegistration) (*AuthPayload, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) LoginByPassword(ctx context.Context, input *LoginByPassword) (*AuthPayload, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Registry(ctx context.Context, input *Registry) (*wrapperspb.BoolValue, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) App(ctx context.Context, appID string) (*AppInfo, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) VideoDownload(ctx context.Context, url string) (*VideoInfo, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) User(ctx context.Context) (*UserInformation, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Captcha(ctx context.Context) (*Captcha, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
