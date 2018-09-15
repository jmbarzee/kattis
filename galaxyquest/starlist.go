package main

type (
	Star struct {
		X    int64
		Y    int64
		Next *Star
		Prev *Star
	}

	StarList struct {
		Head *Star
		Tail *Star
		Len  int
	}
)

func NewStarList() *StarList {
	return &StarList{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

func (sl *StarList) Remove(star *Star) {
	sl.Len--
	if sl.Head == star {
		// Star is at the Front
		star.Next.Prev = nil
		sl.Head = star.Next
	} else if sl.Tail == star {
		// Star is at the Back
		star.Prev.Next = nil
		sl.Tail = star.Prev
	} else {
		// Star is in the Middle
		star.Prev.Next = star.Next
		star.Next.Prev = star.Prev
	}
}

func (sl *StarList) Append(star *Star) {
	sl.Len++
	if sl.Len == 1 {
		sl.Head = star
		sl.Tail = star
	}
	sl.Tail.Next = star
	star.Prev = sl.Tail
	sl.Tail = star
}

func SameGalaxy(s1, s2 *Star, distance int64) bool {
	xDiff := s1.X - s2.X
	x2 := xDiff * xDiff
	yDiff := s1.Y - s2.Y
	y2 := yDiff * yDiff
	d2 := distance * distance
	return (x2 + y2) <= d2
}
