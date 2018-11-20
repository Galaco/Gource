package cache

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	entity2 "github.com/galaco/Gource-Engine/entity"
)

type PropCache struct {
	cacheList []Entry
}

func (c *PropCache) NeedsRecache() bool {
	if len(c.cacheList) == 0 {
		return true
	}
	return false
}

func (c *PropCache) Add(props ...entity.IEntity) {
	for _, prop := range props {
		c.cacheList = append(c.cacheList, Entry{
			Transform: prop.Transform(),
			Model:     prop.(entity2.IProp).GetModel(),
		})
	}
}

func (c *PropCache) All() *[]Entry {
	return &c.cacheList
}

type Entry struct {
	Transform *entity.Transform
	Model     *model.Model
}

func NewPropCache(props ...entity.IEntity) *PropCache {
	c := &PropCache{}

	if len(props) > 0 {
		c.Add(props...)
	}

	return c
}