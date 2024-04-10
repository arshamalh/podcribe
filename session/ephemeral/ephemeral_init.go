package ephemeral

import "podcribe/session"

type ephemeral struct {
	userData map[int64]*UserData
}

func New() *ephemeral {
	return &ephemeral{
		userData: make(map[int64]*UserData),
	}
}

func (e *ephemeral) GetClient(userID int64) session.UserDataSession {
	_, ok := e.userData[userID]
	if !ok {
		e.init(userID)
	}
	return e.userData[userID]
}

func (e *ephemeral) init(userID int64) {
	e.userData[userID] = &UserData{
		id: userID,
	}
}
