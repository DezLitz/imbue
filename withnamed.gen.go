// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// With0Named describes how to construct named values of type T.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With0Named[N Name[T], T any](
	con *Container,
	ctor func(Context) (T, error),
	options ...WithNamedOption,
) {
	With0(
		con,
		func(ctx Context) (ByName[N, T], error) {
			v, err := ctor(ctx)
			return withName[N](v), err
		},
	)
}

// With1Named describes how to construct named values of type T from a single
// dependency.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With1Named[N Name[T], T, D any](
	con *Container,
	ctor func(Context, D) (T, error),
	options ...WithNamedOption,
) {
	With1(
		con,
		func(ctx Context, v1 D) (ByName[N, T], error) {
			v, err := ctor(ctx, v1)
			return withName[N](v), err
		},
	)
}

// With2Named describes how to construct named values of type T from 2
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With2Named[N Name[T], T, D1, D2 any](
	con *Container,
	ctor func(Context, D1, D2) (T, error),
	options ...WithNamedOption,
) {
	With2(
		con,
		func(ctx Context, v1 D1, v2 D2) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2)
			return withName[N](v), err
		},
	)
}

// With3Named describes how to construct named values of type T from 3
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With3Named[N Name[T], T, D1, D2, D3 any](
	con *Container,
	ctor func(Context, D1, D2, D3) (T, error),
	options ...WithNamedOption,
) {
	With3(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3)
			return withName[N](v), err
		},
	)
}

// With4Named describes how to construct named values of type T from 4
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With4Named[N Name[T], T, D1, D2, D3, D4 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4) (T, error),
	options ...WithNamedOption,
) {
	With4(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3, v4 D4) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4)
			return withName[N](v), err
		},
	)
}

// With5Named describes how to construct named values of type T from 5
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With5Named[N Name[T], T, D1, D2, D3, D4, D5 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5) (T, error),
	options ...WithNamedOption,
) {
	With5(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5)
			return withName[N](v), err
		},
	)
}

// With6Named describes how to construct named values of type T from 6
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With6Named[N Name[T], T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6) (T, error),
	options ...WithNamedOption,
) {
	With6(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6)
			return withName[N](v), err
		},
	)
}

// With7Named describes how to construct named values of type T from 7
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With7Named[N Name[T], T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...WithNamedOption,
) {
	With7(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6, v7 D7) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6, v7)
			return withName[N](v), err
		},
	)
}

// With8Named describes how to construct named values of type T from 8
// dependencies.
//
// N is the name given to the dependency.
// T is the type of the dependency.
func With8Named[N Name[T], T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	ctor func(Context, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...WithNamedOption,
) {
	With8(
		con,
		func(ctx Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6, v7 D7, v8 D8) (ByName[N, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6, v7, v8)
			return withName[N](v), err
		},
	)
}
