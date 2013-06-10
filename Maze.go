package main
import "bytes"

type MazeI interface {
    RuneAt(Position) (rune)
    Width() int
    Height() int
}

func ToString(m MazeI) string {
    var buffer bytes.Buffer
    width := m.Width()
    height:= m.Height()
    var p Position
    for p.row = 0; p.row < height; p.row++ {
        for p.col = 0; p.col < width; p.col++ {
            buffer.WriteRune(m.RuneAt(p))
        }
        buffer.WriteByte('\n')
    }
    return buffer.String()
}

