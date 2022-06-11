package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// generateInvokeWithFunc generates the InvokeWithX function for the given
// number of dependencies.
func generateInvokeWithFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("InvokeWith%d", depCount)

	switch depCount {
	case 1:
		code.
			Commentf(
				"%s calls a function with a single dependency.",
				name,
			)
	default:
		code.
			Commentf(
				"%s calls a function with %d dependencies.",
				name,
				depCount,
			)
	}

	impl := &jen.Statement{}

	for n := 0; n < depCount; n++ {
		impl.Add(
			jen.List(
				jen.Id(dependencyParamName(depCount, n)),
				jen.Err(),
			).
				Op(":=").
				Qual(pkgPath, "get").
				Index(
					jen.Id(dependencyTypeName(depCount, n)),
				).
				Call(
					jen.Qual(pkgPath, "rootContext").
						Call(
							contextName(),
						),
					containerName(),
				),
		)

		impl.Add(
			jen.If(
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(
					jen.Err(),
				),
			),
		)
	}

	impl.Add(
		jen.Return(
			jen.Id("fn").Call(
				inputParamNames(depCount)...,
			),
		),
	)

	code.
		Func().
		Id(name).
		Types(typeParams(false, depCount)...).
		Params(
			jen.Line().
				Add(stdContextParam()),
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputParamTypes(stdContextType(), depCount)...,
				).
				Params(
					jen.Error(),
				),
			jen.Line(),
		).
		Params(
			jen.Error(),
		).
		Block(*impl...)
}
