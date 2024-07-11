package ephemeral

import "podcribe/entities"

type UserData struct {
	id    int64
	scene entities.Scene
}

func (ud *UserData) ID() int64 {
	return ud.id
}

func (d *UserData) SetScene(scene entities.Scene) {
	d.scene = scene
}

func (d *UserData) GetScene() entities.Scene {
	return d.scene
}
