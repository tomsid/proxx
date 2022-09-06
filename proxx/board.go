package proxx

// The first idea was to use a simple two-dimensional slice, but then I realized it'd be hard to do
// the traversal of adjacent cells. That's why a graph approach was chosen, where every cell knows about near located cells
type board [][]*cell
