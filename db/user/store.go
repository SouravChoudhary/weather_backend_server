/*
In this implementation:
We define a userRepository struct that contains a reference to the GORM database (db).
The NewUserRepository function creates and returns a new UserRepository instance.
The CreateUser function inserts a new user record into the database.
The GetUserByID function retrieves a user by their ID.
The GetUserByUsername function retrieves a user by their username.
The UpdateUser function updates a user's information in the database.
The DeleteUser function deletes a user by their ID.
*/

package user

import (
	"weatherapp/apperror"
	"weatherapp/pkg/logging"

	"gorm.io/gorm"
)

type Store interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type storeGorm struct {
	db     *gorm.DB
	logger logging.Logger
}

func NewStore(db *gorm.DB, logger logging.Logger) Store {
	return &storeGorm{db: db, logger: logger}
}

func (us *storeGorm) Create(user *User) error {
	if err := us.db.Create(user).Error; err != nil {
		us.logger.LogError("fail to create user", map[string]interface{}{"package": "user", "method": "Create", "user_id": user.ID, "error": err.Error()})
		return apperror.ReturnStoreErr("U04", err)
	}
	us.logger.LogInfo("user created successfully", map[string]interface{}{"user_id": user.ID})
	return nil
}

func (us *storeGorm) GetByID(id int) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.UserNotFound
		}
		us.logger.LogError("fail to get user by id", map[string]interface{}{"package": "user", "method": "GetByID", "user_id": id, "error": err.Error()})
		return nil, apperror.ReturnStoreErr("U04", err)
	}
	return &user, nil
}

func (us *storeGorm) GetByUsername(username string) (*User, error) {
	var user User
	err := us.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.UserNotFound
		}
		us.logger.LogError("fail to get user by username", map[string]interface{}{"package": "user", "method": "GetByUsername", "user_name": username, "error": err.Error()})
		return nil, apperror.ReturnStoreErr("U04", err)
	}
	return &user, nil
}

func (us *storeGorm) Update(user *User) error {
	if err := us.db.Save(user).Error; err != nil {
		us.logger.LogError("fail to update user", map[string]interface{}{"package": "user", "method": "Update", "user_id": user.ID, "error": err.Error()})
		return apperror.ReturnStoreErr("U05", err)
	}
	us.logger.LogInfo("user updated successfully", map[string]interface{}{"user_id": user.ID})
	return nil
}

func (us *storeGorm) Delete(id uint) error {
	if err := us.db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		us.logger.LogError("fail to delete user", map[string]interface{}{"package": "user", "method": "Delete", "user_id": id, "error": err.Error()})
		return apperror.ReturnStoreErr("U03", err)
	}
	us.logger.LogInfo("user deleted successfully", map[string]interface{}{"user_id": id})
	return nil
}
