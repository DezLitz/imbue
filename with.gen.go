// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// With0 describes how to construct values of type T.
func With0[T any](
	con *Container,
	ctor func(Context) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			return ctor(ctx)
		},
	)
}

// With1 describes how to construct values of type T from a single dependency.
func With1[T, D any](
	con *Container,
	ctor func(Context, D) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1)
		},
		d1,
	)
}

// With2 describes how to construct values of type T from 2 dependencies.
func With2[T, D1, D2 any](
	con *Container,
	ctor func(Context, D1, D2) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2)
		},
		d1,
		d2,
	)
}

// With3 describes how to construct values of type T from 3 dependencies.
func With3[T, D1, D2, D3 any](
	con *Container,
	ctor func(Context, D1, D2, D3) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3)
		},
		d1,
		d2,
		d3,
	)
}

// With4 describes how to construct values of type T from 4 dependencies.
func With4[T, D1, D2, D3, D4 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)
	d4 := get[D4](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v4, err := d4.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3, v4)
		},
		d1,
		d2,
		d3,
		d4,
	)
}

// With5 describes how to construct values of type T from 5 dependencies.
func With5[T, D1, D2, D3, D4, D5 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)
	d4 := get[D4](con)
	d5 := get[D5](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v4, err := d4.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v5, err := d5.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3, v4, v5)
		},
		d1,
		d2,
		d3,
		d4,
		d5,
	)
}

// With6 describes how to construct values of type T from 6 dependencies.
func With6[T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)
	d4 := get[D4](con)
	d5 := get[D5](con)
	d6 := get[D6](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v4, err := d4.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v5, err := d5.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v6, err := d6.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3, v4, v5, v6)
		},
		d1,
		d2,
		d3,
		d4,
		d5,
		d6,
	)
}

// With7 describes how to construct values of type T from 7 dependencies.
func With7[T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)
	d4 := get[D4](con)
	d5 := get[D5](con)
	d6 := get[D6](con)
	d7 := get[D7](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v4, err := d4.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v5, err := d5.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v6, err := d6.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v7, err := d7.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3, v4, v5, v6, v7)
		},
		d1,
		d2,
		d3,
		d4,
		d5,
		d6,
		d7,
	)
}

// With8 describes how to construct values of type T from 8 dependencies.
func With8[T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...WithOption,
) {
	t := get[T](con)
	d1 := get[D1](con)
	d2 := get[D2](con)
	d3 := get[D3](con)
	d4 := get[D4](con)
	d5 := get[D5](con)
	d6 := get[D6](con)
	d7 := get[D7](con)
	d8 := get[D8](con)

	t.Declare(
		func(ctx Context) (v T, _ error) {
			v1, err := d1.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v2, err := d2.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v3, err := d3.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v4, err := d4.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v5, err := d5.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v6, err := d6.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v7, err := d7.Resolve(ctx)
			if err != nil {
				return v, err
			}

			v8, err := d8.Resolve(ctx)
			if err != nil {
				return v, err
			}

			return ctor(ctx, v1, v2, v3, v4, v5, v6, v7, v8)
		},
		d1,
		d2,
		d3,
		d4,
		d5,
		d6,
		d7,
		d8,
	)
}
