//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package golang

import (
	"github.com/goccmack/gogll/ast"
	"github.com/goccmack/gogll/cfg"
	"github.com/goccmack/gogll/frstflw"
	"github.com/goccmack/gogll/gen/golang/goutil"
	"github.com/goccmack/gogll/gen/golang/parser"
	"github.com/goccmack/gogll/gslot"
	"path"
)

func Gen(g *ast.Grammar, gs *gslot.GSlot, ff *frstflw.FF) {
	goutil.Gen(path.Join(cfg.BaseDir, "goutil"), g)
	parser.Gen(path.Join(cfg.BaseDir, "parser"), g, gs, ff)
}
