// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

import "context"

// Invoke1 calls a function with a single dependency.
func Invoke1[D any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1)
}

// Invoke2 calls a function with 2 dependencies.
func Invoke2[D1, D2 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2)
}

// Invoke3 calls a function with 3 dependencies.
func Invoke3[D1, D2, D3 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3)
}

// Invoke4 calls a function with 4 dependencies.
func Invoke4[D1, D2, D3, D4 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	d4, err := get[D4](con)
	if err != nil {
		return err
	}

	v4, err := d4.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3, v4)
}

// Invoke5 calls a function with 5 dependencies.
func Invoke5[D1, D2, D3, D4, D5 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	d4, err := get[D4](con)
	if err != nil {
		return err
	}

	v4, err := d4.Resolve(rctx)
	if err != nil {
		return err
	}

	d5, err := get[D5](con)
	if err != nil {
		return err
	}

	v5, err := d5.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3, v4, v5)
}

// Invoke6 calls a function with 6 dependencies.
func Invoke6[D1, D2, D3, D4, D5, D6 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	d4, err := get[D4](con)
	if err != nil {
		return err
	}

	v4, err := d4.Resolve(rctx)
	if err != nil {
		return err
	}

	d5, err := get[D5](con)
	if err != nil {
		return err
	}

	v5, err := d5.Resolve(rctx)
	if err != nil {
		return err
	}

	d6, err := get[D6](con)
	if err != nil {
		return err
	}

	v6, err := d6.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3, v4, v5, v6)
}

// Invoke7 calls a function with 7 dependencies.
func Invoke7[D1, D2, D3, D4, D5, D6, D7 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	d4, err := get[D4](con)
	if err != nil {
		return err
	}

	v4, err := d4.Resolve(rctx)
	if err != nil {
		return err
	}

	d5, err := get[D5](con)
	if err != nil {
		return err
	}

	v5, err := d5.Resolve(rctx)
	if err != nil {
		return err
	}

	d6, err := get[D6](con)
	if err != nil {
		return err
	}

	v6, err := d6.Resolve(rctx)
	if err != nil {
		return err
	}

	d7, err := get[D7](con)
	if err != nil {
		return err
	}

	v7, err := d7.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3, v4, v5, v6, v7)
}

// Invoke8 calls a function with 8 dependencies.
func Invoke8[D1, D2, D3, D4, D5, D6, D7, D8 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8) error,
	options ...InvokeOption,
) error {
	rctx := rootContext(ctx, con)

	d1, err := get[D1](con)
	if err != nil {
		return err
	}

	v1, err := d1.Resolve(rctx)
	if err != nil {
		return err
	}

	d2, err := get[D2](con)
	if err != nil {
		return err
	}

	v2, err := d2.Resolve(rctx)
	if err != nil {
		return err
	}

	d3, err := get[D3](con)
	if err != nil {
		return err
	}

	v3, err := d3.Resolve(rctx)
	if err != nil {
		return err
	}

	d4, err := get[D4](con)
	if err != nil {
		return err
	}

	v4, err := d4.Resolve(rctx)
	if err != nil {
		return err
	}

	d5, err := get[D5](con)
	if err != nil {
		return err
	}

	v5, err := d5.Resolve(rctx)
	if err != nil {
		return err
	}

	d6, err := get[D6](con)
	if err != nil {
		return err
	}

	v6, err := d6.Resolve(rctx)
	if err != nil {
		return err
	}

	d7, err := get[D7](con)
	if err != nil {
		return err
	}

	v7, err := d7.Resolve(rctx)
	if err != nil {
		return err
	}

	d8, err := get[D8](con)
	if err != nil {
		return err
	}

	v8, err := d8.Resolve(rctx)
	if err != nil {
		return err
	}

	return fn(ctx, v1, v2, v3, v4, v5, v6, v7, v8)
}
