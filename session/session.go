package session

// TelegramSession is a temporal+fast storage used for storing user current state
type TelegramSession interface {
	GetClient(userID int64) UserDataSession
}

// UserDataSession is a user-aware session, meaning that it knows you are setting things for which user
type UserDataSession interface {
	ID() int64
}
