package main

type Cell struct {
    rune      rune
    seen      bool
    enterable bool
}

func (cell *Cell) Rune() rune {
    if cell.seen {
        return cell.rune
    }
    return WALL
}
