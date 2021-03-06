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
package goutil

import (
	"github.com/goccmack/gogll/ast"
	"github.com/goccmack/gogll/gen/golang/goutil/bsr"
	"github.com/goccmack/gogll/gen/golang/goutil/md"
	"github.com/goccmack/gogll/gen/golang/goutil/stringset"
	"path/filepath"
)

func Gen(utilDir string, g *ast.Grammar) {
	bsr.Gen(filepath.Join(utilDir, "bsr", "bsr.go"), g.Package)
	stringset.Gen(filepath.Join(utilDir, "stringset", "stringset.go"))
	md.Gen(filepath.Join(utilDir, "md", "md.go"))
}
