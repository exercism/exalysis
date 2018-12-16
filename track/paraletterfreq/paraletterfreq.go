package paraletterfreq

import (
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/track/paraletterfreq/tpl"
)

//Suggest builds suggestions for the exercise solution
func Suggest(pkg *astrav.Package, r *extypes.Response) {
	for _, tf := range exFuncs {
		tf(pkg, r)
	}
}

var exFuncs = []extypes.SuggestionFunc{
	addConcurrencyNotFaster,
	examAddOne,
	examWaitGroup,
	examBufferSizeLen,
	examSelect,
	examGoroutineLeak,
	examForRange,
}

func addConcurrencyNotFaster(_ *astrav.Package, r *extypes.Response) {
	r.AppendOutro(tpl.ConcurrencyNotFaster)
}

func examWaitGroup(pkg *astrav.Package, r *extypes.Response) {
	wgs := pkg.FindByName("WaitGroup")
	goTokens := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(wgs) != 0 && 1 < len(goTokens) {
		r.AppendImprovement(tpl.WaitGroup)
	}
	if len(wgs) != 0 && len(goTokens) == 1 {
		r.AppendImprovement(tpl.WaitGroupNotNeeded)
	}
}

func examAddOne(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("Add")
	for _, node := range nodes {
		bLit := node.Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if bLit == nil {
			continue
		}
		if bLit.(*astrav.BasicLit).Value == "1" {
			r.AppendImprovement(tpl.WaitGroupAddOne)
		}
	}
}

func examBufferSizeLen(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("make")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeCallExpr) {
			continue
		}
		lenNode := node.FindFirstByName("len")
		if lenNode != nil {
			r.AppendImprovement(tpl.BufferSizeLen)
		}
	}
}

func examSelect(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeSelectStmt)
	for _, node := range nodes {
		if len(node.Children()) == 1 {
			r.AppendImprovement(tpl.SelectNotNeeded)
		}
	}
}

func examGoroutineLeak(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	for _, node := range nodes {
		loops := node.FindByNodeType(astrav.NodeTypeForStmt)
		for _, loop := range loops {
			cond := loop.(*astrav.ForStmt).Cond()
			if cond == nil {
				r.AppendTodo(tpl.GoroutineLeak)
			}
		}
	}
}

func examForRange(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	if len(nodes) == 0 {
		return
	}
	mergeLoop := nodes[len(nodes)-1]
	cond := mergeLoop.(*astrav.ForStmt).Cond()
	if cond != nil {
		r.AppendImprovement(tpl.ForRangeNoVars)
	}
}
