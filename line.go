package main

import "math"

// Line is a struct of a finite line between points a and b
type Line struct {
	A *Vector `json:"a"`
	B *Vector `json:"b"`
}

// Interect returns the point of intersection another line
func (line *Line) OldIntersect(line2 Line) []*Vector {
	var intersections []*Vector
	d := (line.A.X-line.B.X)*(line2.A.Y-line2.B.Y) - (line.A.Y-line.B.Y)*(line2.A.X-line2.B.X)
	if d != 0 {
		xi := ((line2.A.X-line2.B.X)*(line.A.X*line.B.Y-line.A.Y*line.B.X) -
			(line.A.X-line.B.X)*(line2.A.X*line2.B.Y-line2.A.Y*line2.B.X)) / d
		yi := ((line2.A.Y-line2.B.Y)*(line.A.X*line.B.Y-line.A.Y*line.B.X) -
			(line.A.Y-line.B.Y)*(line2.A.X*line2.B.Y-line2.A.Y*line2.B.X)) / d
		if !((line.A.X < xi && line.B.X < xi) || (line.A.X > xi && line.B.X > xi) ||
			(line.A.Y < yi && line.B.Y < yi) || (line.A.Y > yi && line.B.Y > yi) ||
			(line2.A.X < xi && line2.B.X < xi) || (line2.A.X > xi && line2.B.X > xi) ||
			(line2.A.Y < yi && line2.B.Y < yi) || (line2.A.Y > yi && line2.B.Y > yi)) {
			intersections = append(intersections, &Vector{xi, yi})
		}
	}
	return intersections
}

func (a *Line) Intersect(b Line) []*Vector {
	Ax := a.A.X
	Ay := a.A.Y
	Bx := a.B.X
	By := a.B.Y
	Cx := b.A.X
	Cy := b.A.Y
	Dx := b.B.X
	Dy := b.B.Y

	// Zero length
	if Ax == Bx && Ay == By || Cx == Dx && Cy == Dy {
		return nil
	}

	// Shares endpoint
	if Ax == Cx && Ay == Cy || Bx == Cx && By == Cy || Ax == Dx && Ay == Dy || Bx == Dx && By == Dy {
		return nil
	}

	// Translate so A = origin
	Bx -= Ax
	By -= Ay
	Cx -= Ax
	Cy -= Ay
	Dx -= Ax
	Dy -= Ay

	distAB := math.Sqrt(Bx*Bx + By*By)

	theCos := Bx / distAB
	theSin := By / distAB
	newX := Cx*theCos + Cy*theSin
	Cy = Cy*theCos - Cx*theSin
	Cx = newX
	newX = Dx*theCos + Dy*theSin
	Dy = Dy*theCos - Dx*theSin
	Dx = newX

	// A-B doesn't cross C-D
	if Cy < 0 && Dy < 0 || Cy >= 0 && Dy >= 0 {
		return nil
	}

	ABpos := Dx + (Cx-Dx)*Dy/(Dy-Cy)

	// A-B crosses C-D outside
	if ABpos < 0 || ABpos > distAB {
		return nil
	}

	X := Ax + ABpos*theCos
	Y := Ay + ABpos*theSin

	return []*Vector{&Vector{X, Y}}
}
