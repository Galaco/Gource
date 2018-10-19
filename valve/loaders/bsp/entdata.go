package bsp

import (
	"github.com/galaco/Gource-Engine/components"
	"github.com/galaco/Gource-Engine/components/renderable"
	"github.com/galaco/Gource-Engine/engine/base/primitive"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/interfaces"
	entity2 "github.com/galaco/Gource-Engine/entity"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/vmf"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

// Parse Entity block.
// Vmf lib is actually capable of doing this;
// contents are loaded into Vmf.Unclassified
func ParseEntities(data string) (vmf.Vmf, error) {
	stringReader := strings.NewReader(data)
	reader := vmf.NewReader(stringReader)

	return reader.Read()
}

func CreateEntity(ent *entity.Entity) interfaces.IEntity{
	localEdict := &entity2.ValveEntity{}
	origin := ent.VectorForKey("origin")
	localEdict.GetTransformComponent().Position = mgl32.Vec3{origin.X(), origin.Y(), origin.Z()}
	localEdict.GetTransformComponent().Scale = mgl32.Vec3{8, 8, 8}

	placeholder := components.NewRenderableComponent()
	resource := renderable.NewGPUResource([]interfaces.IPrimitive{primitive.NewCube()})
	resource.Prepare()
	placeholder.AddRenderableResource(resource)
	e := factory.NewEntity(localEdict)
	factory.NewComponent(placeholder, e)

	return e
}
