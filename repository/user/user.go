package user

import (
	"backend-ewallet/entities"
	"backend-ewallet/utils"
	"errors"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {

	user.Password, _ = utils.HashPassword(user.Password)
	uid := shortuuid.New()
	user.UserID = uid

	if err := ur.database.Create(&user).Error; err != nil {
		return user, errors.New("invalid input or this email was created (duplicated entry)")
	}

	return user, nil
}

func (ur *UserRepository) GetByID(userID string) (entities.User, error) {
	arrUser := entities.User{}

	result := ur.database.Where("user_id =?", userID).First(&arrUser)
	if err := result.Error; err != nil {
		return arrUser, err
	}
	if result.RowsAffected == 0 {
		return arrUser, errors.New("record not found")
	}

	return arrUser, nil
}

func (ur *UserRepository) Update(userUid string, newUser entities.User) (entities.User, error) {

	var user entities.User
	result := ur.database.Where("user_id =?", userUid).First(&user)

	if result.Error != nil {
		return entities.User{}, errors.New("failed to update user")
	}
	if result.RowsAffected == 0 {
		return entities.User{}, errors.New("user not found")
	}

	if err := ur.database.Model(&user).Where("user_id =?", userUid).Updates(&newUser).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userUid string) error {

	result := ur.database.Where("user_id =?", userUid).Delete(&entities.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (ur *UserRepository) Search(q string) ([]entities.User, error) {
	arrUser := []entities.User{}

	if len(q) < 3 {
		if len(q) == 1 {
			ur.database.Debug().Where("a =?", q).Find(&arrUser)
			return arrUser, nil
		}
		if len(q) == 2 {
			ur.database.Debug().Where("b =?", q).Find(&arrUser)
			return arrUser, nil

		}
		if len(q) == 3 {
			ur.database.Debug().Where("c =?", q).Find(&arrUser)
			return arrUser, nil

		}
	}
	sql := "%" + q + "%"

	result := ur.database.Debug().Where("name like ?", sql).Find(&arrUser)
	if err := result.Error; err != nil {
		return arrUser, err
	}
	if result.RowsAffected == 0 {
		return arrUser, errors.New("record not found")
	}

	return arrUser, nil
}

// func (ur *UserRepository) Dummy(length int) bool {

// 	names := []string{"roger", "joni", "mail", "bruto", "icon", "abeng", "jangkrik", "zeagger", "connie", "terlalu"}

// 	for i := 0; i < length; i++ {
// 		uid := shortuuid.New()

// 		user := entities.User{
// 			UserUid: uid,
// 			Name:    names[rand.Intn(9)],
// 			Email:   faker.Email(),
// 			Address: "jl.dramaga no.22",
// 			Gender:  "male",
// 		}
// 		user.Password = "xyz"
// 		user.Password, _ = middlewares.HashPassword(user.Password)

// 		if err := ur.database.Create(&user).Error; err != nil {
// 			return false
// 		}
// 	}

// 	return true
// }
