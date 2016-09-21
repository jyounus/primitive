package primitive

import (
	"image"
	"math/rand"
)

type State struct {
	Model  *Model
	Buffer *image.RGBA
	Shape  Shape
	Alpha  int
}

func NewState(model *Model, buffer *image.RGBA, shape Shape) *State {
	return &State{model, buffer, shape, 128}
}

func (state *State) Energy() float64 {
	return state.Model.Energy(state.Shape, state.Alpha, state.Buffer)
}

func (state *State) DoMove() interface{} {
	oldShape := state.Shape.Copy()
	oldAlpha := state.Alpha
	state.Shape.Mutate()
	state.Alpha = clampInt(state.Alpha+rand.Intn(21)-10, 0, 255)
	return State{nil, nil, oldShape, oldAlpha}
}

func (state *State) UndoMove(undo interface{}) {
	oldState := undo.(State)
	state.Shape = oldState.Shape
	state.Alpha = oldState.Alpha
}

func (state *State) Copy() Annealable {
	return &State{state.Model, state.Buffer, state.Shape.Copy(), state.Alpha}
}
