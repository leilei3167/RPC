package service

import "context"

type R struct {
	S string
}
type In struct {
	I string
}

type A struct{}
type B struct{}

func (t *A) Addhello(ctx context.Context, in *In, r *R) error {
	r.S = in.I + "helloworld"

	return nil
}

func (t *A) Addgin(ctx context.Context, in *In, r *R) error {
	r.S = in.I + "gin"

	return nil
}

func (t *B) Add1(ctx context.Context, in *In, r *R) error {
	r.S = in.I + "1111111111111"

	return nil
}

func (t *B) Add2(ctx context.Context, in *In, r *R) error {
	r.S = in.I + "222222222"

	return nil
}
