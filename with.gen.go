// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// With0 describes how to construct values of type T.
func With0[T any](
	con *Container,
	fn func(*Context) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			return fn, nil
		},
	); err != nil {
		panic(err)
	}
}

// With1 describes how to construct values of type T from a single dependency.
func With1[T, D any](
	con *Container,
	fn func(*Context, D) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return fn(ctx, v1)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With2 describes how to construct values of type T from 2 dependencies.
func With2[T, D1, D2 any](
	con *Container,
	fn func(*Context, D1, D2) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return fn(ctx, v1, v2)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With3 describes how to construct values of type T from 3 dependencies.
func With3[T, D1, D2, D3 any](
	con *Container,
	fn func(*Context, D1, D2, D3) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With4 describes how to construct values of type T from 4 dependencies.
func With4[T, D1, D2, D3, D4 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddConstructorDependency(d4); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3, v4)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With5 describes how to construct values of type T from 5 dependencies.
func With5[T, D1, D2, D3, D4, D5 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddConstructorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddConstructorDependency(d5); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3, v4, v5)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With6 describes how to construct values of type T from 6 dependencies.
func With6[T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddConstructorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddConstructorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddConstructorDependency(d6); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3, v4, v5, v6)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With7 describes how to construct values of type T from 7 dependencies.
func With7[T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddConstructorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddConstructorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddConstructorDependency(d6); err != nil {
				return nil, err
			}

			d7 := get[D7](con)
			if err := t.AddConstructorDependency(d7); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3, v4, v5, v6, v7)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// With8 describes how to construct values of type T from 8 dependencies.
func With8[T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...WithOption,
) {
	t := get[T](con)

	if err := t.Declare(
		func() (constructor[T], error) {
			d1 := get[D1](con)
			if err := t.AddConstructorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddConstructorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddConstructorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddConstructorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddConstructorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddConstructorDependency(d6); err != nil {
				return nil, err
			}

			d7 := get[D7](con)
			if err := t.AddConstructorDependency(d7); err != nil {
				return nil, err
			}

			d8 := get[D8](con)
			if err := t.AddConstructorDependency(d8); err != nil {
				return nil, err
			}

			return func(ctx *Context) (v T, _ error) {
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

				return fn(ctx, v1, v2, v3, v4, v5, v6, v7, v8)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}
