package store

type Member struct {
	// PK
	ID uint `gorm:"primary_key"`

	Name string
}

// 新增會員
func (db *dataStore) CreateMember(member *Member) error {
	return db.Create(member).Error
}
