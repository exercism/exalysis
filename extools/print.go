package extools

import (
	"fmt"

	"github.com/tehsphinx/dbg"
	"golang.org/x/tools/go/loader"
)

//PrintAST prints useful information on a package
func PrintAST(pkg *loader.PackageInfo) {
	PrintTypes(pkg)
	PrintUses(pkg)
	PrintScopes(pkg)
	PrintImplicits(pkg)
	PrintDefs(pkg)
	PrintSelections(pkg)
}

//PrintUses prints all uses of given package
func PrintUses(pkg *loader.PackageInfo) {
	printHead("Uses")
	for k, v := range pkg.Uses {
		fmt.Println("name:", k, "\t\tvalue:", v)
	}
}

//PrintScopes prints all scopes of given package
func PrintScopes(pkg *loader.PackageInfo) {
	printHead("Scopes")
	for _, v := range pkg.Scopes {
		fmt.Println(v)
	}
}

//PrintImplicits prints all implicits of given package
func PrintImplicits(pkg *loader.PackageInfo) {
	printHead("Implicits")
	for k, v := range pkg.Implicits {
		fmt.Println("name:", k, "\t\tvalue:", v)
	}
}

//PrintDefs prints all definitions of given package
func PrintDefs(pkg *loader.PackageInfo) {
	printHead("Definitions")
	for k, v := range pkg.Defs {
		fmt.Println("name:", k, "\t\tvalue:", v)
	}
}

//PrintSelections prints all selections of given package
func PrintSelections(pkg *loader.PackageInfo) {
	printHead("Selections")
	for k, v := range pkg.Selections {
		fmt.Println("name:", k, "\t\tvalue:", v)
	}
}

//PrintTypes prints all types of given package
func PrintTypes(pkg *loader.PackageInfo) {
	printHead("Types")
	for k, v := range pkg.Types {
		fmt.Println("name:", k, "\t\tvalue:", v)
	}
}

func printHead(name string) {
	dbg.Blue("################")
	dbg.Blue("#### " + name)
	dbg.Blue("################")

}
