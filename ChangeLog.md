# Change Log

# 2019-09-13
1. Documentation added on how to walk the parse forest. [See](doc/bsr/bsr.md)

## 2019-09-10
1. testSelect is generated as a slice of functions, not a map of functions.
1. follow is generated as separate functions, not a map of functions.

## 2019-09-10 v2.0.0
1. Bug fixed in `ast.NewAnyOf` and `ast.NewNot`
1. Grammar change: `emptyAlt` replace by `empty`. `emptyAlt` was a vestige of 
the bootstrap gocc compiler.

## 2019-09-09
BSR performance improved. 