package main

import (
	"github.com/rzamoramx/test-go-modules/submodule"
	v2 "github.com/rzamoramx/test-go-modules/submodule/v2"

	"github.com/rzamoramx/test-go-modules/othermodule"
	omv2 "github.com/rzamoramx/test-go-modules/othermodule/v2"

	"bitbucket.org/cargamos/saas-apiv3-lib-temp/othermod"
	x "bitbucket.org/cargamos/saas-apiv3-lib-temp/submodule"
	xv2 "bitbucket.org/cargamos/saas-apiv3-lib-temp/submodule/v2"
)

func main() {
	submodule.Foo()
	v2.Foo()

	othermodule.Bar()
	omv2.Bar()

	x.Foo()
	xv2.Foo()

	othermod.Bar()
}
