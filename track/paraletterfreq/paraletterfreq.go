package paraletterfreq

import (
	"strings"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/track/paraletterfreq/tpl"
	"github.com/tehsphinx/astrav"
)

// Suggest builds suggestions for the exercise solution
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
	examCombineWhileWaiting,
	examMutex,
	examRangeChan,
}

func addConcurrencyNotFaster(_ *astrav.Package, r *extypes.Response) {
	r.AppendOutro(tpl.ConcurrencyNotFaster)
}

func examWaitGroup(pkg *astrav.Package, r *extypes.Response) {
	wgs := pkg.FindByName("WaitGroup")
	goTokens := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(wgs) != 0 && 1 < len(goTokens) {
		r.AppendImprovementTpl(tpl.WaitGroup)
	}
	if len(wgs) != 0 && len(goTokens) == 1 {
		r.AppendImprovementTpl(tpl.WaitGroupNotNeeded)
	}
}

func examAddOne(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByName("Add")
	for _, node := range nodes {
		bLit := node.NextParentByType(astrav.NodeTypeCallExpr).FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if bLit == nil {
			continue
		}
		if bLit.(*astrav.BasicLit).Value == "1" {
			r.AppendImprovementTpl(tpl.WaitGroupAddOne)
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
			r.AppendImprovementTpl(tpl.BufferSizeLen)
		}
	}
}

func examSelect(pkg *astrav.Package, r *extypes.Response) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeSelectStmt)
	for _, node := range nodes {
		if len(node.Children()) == 1 {
			r.AppendImprovementTpl(tpl.SelectNotNeeded)
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
				r.AppendTodoTpl(tpl.GoroutineLeak)
			}
		}
	}
}

func examForRange(pkg *astrav.Package, r *extypes.Response) {
	loops := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	if len(loops) == 0 {
		return
	}
	mergeLoop := loops[len(loops)-1]
	cond := mergeLoop.(*astrav.ForStmt).Cond()
	if cond != nil {
		r.AppendImprovementTpl(tpl.ForRangeNoVars)
	}
}

func examCombineWhileWaiting(pkg *astrav.Package, r *extypes.Response) {
	loops := pkg.FindByNodeType(astrav.NodeTypeRangeStmt)
	if len(loops) == 0 {
		return
	}
	mergeLoop := loops[len(loops)-1]
	for _, node := range mergeLoop.FindByNodeType(astrav.NodeTypeAssignStmt) {
		token := node.(*astrav.AssignStmt).Tok.String()
		if token == "+=" {
			return
		}
	}
	// They didn't update the map inside the merge loop
	r.AppendImprovementTpl(tpl.CombineMapsWhileWaiting)
}

func examMutex(pkg *astrav.Package, r *extypes.Response) {
	mutexes := pkg.FindByValueType("sync.Mutex")
	if len(mutexes) > 0 {
		r.AppendImprovementTpl(tpl.Mutex)
	}
}

func examRangeChan(pkg *astrav.Package, r *extypes.Response) {
	for _, node := range pkg.FindByNodeType(astrav.NodeTypeRangeStmt) {
		x := node.(*astrav.RangeStmt).X()
		if strings.Contains(x.ValueType().String(), "chan") {
			r.AppendImprovementTpl(tpl.RangeChan)
		}
	}
}
