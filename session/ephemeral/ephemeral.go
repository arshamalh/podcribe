package ephemeral

type UserData struct {
	id int64
}

func (ud *UserData) ID() int64 {
	return ud.id
}
