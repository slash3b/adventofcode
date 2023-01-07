package main

import (
	"fmt"
	"testing"
)

func TestCoord_H_Spot(t *testing.T) {
	parent := &Coord{
		Y: 2,
		X: 0,
	}
	child := &Coord{
		Y: 2,
		X: 2,
	}

	child.Move(parent)

	if child.Y != 2 || child.X != 1 {
		t.Error("failed")
	}
}

func TestCoord_A_Spot(t *testing.T) {
	testCases := []*Coord{
		&Coord{
			Y: 3,
			X: 0,
		},
		&Coord{
			Y: 4,
			X: 0,
		},
		&Coord{
			Y: 4,
			X: 1,
		},
	}

	for _, parent := range testCases {
		child := &Coord{
			Y: 2,
			X: 2,
		}
		child.Move(parent)

		if child.Y != 3 || child.X != 1 {
			t.Error("failed")
		}
	}
}

func TestCoord_B_Spot(t *testing.T) {
	parent := &Coord{
		Y: 4,
		X: 2,
	}
	child := &Coord{
		Y: 2,
		X: 2,
	}

	child.Move(parent)

	if child.Y != 3 || child.X != 2 {
		t.Error("failed")
	}
}

func TestCoord_C_Spot(t *testing.T) {

	testCases := []*Coord{
		&Coord{
			Y: 4,
			X: 3,
		},
		&Coord{
			Y: 4,
			X: 4,
		},
		&Coord{
			Y: 3,
			X: 4,
		},
	}

	for _, parent := range testCases {
		child := &Coord{
			Y: 2,
			X: 2,
		}
		child.Move(parent)

		// fmt.Printf("%#v \n", child)
		if child.Y != 3 || child.X != 3 {
			t.Error("failed")
		}
	}
}

func TestCoord_D_Spot(t *testing.T) {
	parent := &Coord{
		Y: 2,
		X: 4,
	}
	child := &Coord{
		Y: 2,
		X: 2,
	}

	child.Move(parent)

	if child.Y != 2 || child.X != 3 {
		t.Error("failed")
	}
}

func TestCoord_E_Spot(t *testing.T) {

	testCases := []*Coord{
		&Coord{
			Y: 1,
			X: 4,
		},
		&Coord{
			Y: 0,
			X: 4,
		},
		&Coord{
			Y: 0,
			X: 3,
		},
	}

	for _, parent := range testCases {
		child := &Coord{
			Y: 2,
			X: 2,
		}
		child.Move(parent)

		// fmt.Printf("%#v \n", child)
		if child.Y != 1 || child.X != 3 {
			t.Error("failed")
		}
	}
}

func TestCoord_F_Spot(t *testing.T) {
	parent := &Coord{
		Y: 0,
		X: 2,
	}
	child := &Coord{
		Y: 2,
		X: 2,
	}

	child.Move(parent)

	fmt.Printf("%#v \n", child)
	if child.Y != 1 || child.X != 2 {
		t.Error("failed")
	}

}

func TestCoord_G_Spot(t *testing.T) {

	testCases := []*Coord{
		&Coord{
			Y: 0,
			X: 1,
		},
		&Coord{
			Y: 0,
			X: 0,
		},
		&Coord{
			Y: 1,
			X: 0,
		},
	}

	for _, parent := range testCases {
		child := &Coord{
			Y: 2,
			X: 2,
		}
		child.Move(parent)

		fmt.Printf("%#v \n", child)
		if child.Y != 1 || child.X != 1 {
			t.Error("failed")
		}
	}
}
