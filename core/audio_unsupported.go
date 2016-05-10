//+build netgo windows android

package core

import (
	"engo.io/ecs"
)

const (
	defaultHeightModifier float32 = 1
)

var MasterVolume float64 = 1

// ReadSeekCloser is an io.ReadSeeker and io.Closer.
type ReadSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

// AudioComponent is a Component which is used by the AudioSystem
type AudioComponent struct {
	File       string
	Repeat     bool
	Background bool
	RawVolume  float64
}

// AudioSystem is a System that allows for sound effects and / or music
type AudioSystem struct {
	HeightModifier float32
}

func (as *AudioSystem) New(*ecs.World) {
	notImplemented("audio")
}

func (as *AudioSystem) Add(*ecs.BasicEntity, *AudioComponent, *SpaceComponent) {}
func (as *AudioSystem) Remove(basic ecs.BasicEntity)                           {}
func (as *AudioSystem) Update(dt float32)                                      {}