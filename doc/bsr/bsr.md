# Walking the BSR Parser Forest

The parser builds a BSR set in the package `<module>/goutil/bsr` generated by gogll. Package `bsr` exports two functions to walk the parse forest:

1. `func GetRoot() (roots []Slot)` returns the roots of all complete parse trees recognised by the parser. If the grammar is unambiguous there will be only one root.
1. `func (s Slot) GetNTChild(nt string) Slot` returns the BSR `Slot` with matching nonterminal `nt`, left- and right extent.

A parse tree can be walked as follows:

1. Use `GetRoot()` to get the root of the parse tree.
1. Call `GetNTChild(nt)` to get get the BSR of symbol `nt`.