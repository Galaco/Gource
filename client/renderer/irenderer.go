package renderer

import (
	"github.com/galaco/Gource-Engine/client/scene/world"
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderer interface {
	Initialize()
	StartFrame(*entity.Camera)
	LoadShaders()
	DrawBsp(*world.World)
	DrawSkybox(*world.Sky)
	DrawModel(*model.Model, mgl32.Mat4)
	DrawSkyMaterial(*model.Model)
	SetWireframeMode(bool)
	EndFrame()
	Unregister()
}
