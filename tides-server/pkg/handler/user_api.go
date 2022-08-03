package handler

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"

	"tides-server/pkg/restapi/operations/user"

	"tides-server/pkg/config"
	"tides-server/pkg/models"

	"crypto/tls"

	"github.com/sethvargo/go-password/password"
	gomail "gopkg.in/mail.v2"
)

func SendVerificationHandler(params user.SendVerificationParams) middleware.Responder {
	godotenv.Load("/.env")
	OFFICIAL_EMAIL := os.Getenv("OFFICIAL_EMAIL")
	OFFICIAL_PASSWORD := os.Getenv("OFFICIAL_PASSWORD")
	fmt.Println("verification entered!")
	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	fmt.Println(body.Message)
	db := config.GetDB()
	var u models.User
	if db.Where("username = ?", body.Message).First(&u).RowsAffected == 0 {
		return user.NewSendVerificationBadRequest().WithPayload(&user.SendVerificationBadRequestBody{
			Message: "null user",
		})
	}
	code, _ := password.Generate(6, 6, 0, false, false)
	u.Temp = code
	err := db.Save(&u).Error
	if err != nil {
		return user.NewSendVerificationBadRequest().WithPayload(&user.SendVerificationBadRequestBody{
			Message: err.Error(),
		})
	}

	m := gomail.NewMessage()
	m.SetHeader("From", OFFICIAL_EMAIL)
	// m.SetHeader("To", u.Email)
	m.SetHeader("To", u.Email)
	fmt.Println(code)
	m.SetHeader("Subject", "CloudTides Verification Code")
	m.SetBody("text/plain", "Your are resetting your password. Your verification code is: "+code)
	// d := gomail.NewDialer("smtp.gmail.com", 587, OFFICIAL_EMAIL, OFFICIAL_PASSWORD)
	d := gomail.NewDialer("smtp.qiye.aliyun.com", 25, OFFICIAL_EMAIL, OFFICIAL_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	i := 0
	for ; i < 5; i++ {
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	if i == 5 {
		fmt.Println("fail")
		return user.NewSendVerificationBadRequest().WithPayload(&user.SendVerificationBadRequestBody{
			Message: "fail!!!",
		})
	}

	fmt.Println("success!!!")

	return user.NewSendVerificationOK().WithPayload(&user.SendVerificationOKBody{
		Message: "success",
	})

}

// ResetPasswordHandler is the API for /users/reset POST
func ResetPasswordHandler(params user.ResetPasswordParams) middleware.Responder {
	fmt.Println("entered!")
	body := params.ReqBody
	db := config.GetDB()
	var u models.User
	//   fmt.Println(body.Username + ", " + body.Password + ", " + body.NewPassword)
	if db.Where("username = ?", body.Username).First(&u).RowsAffected == 0 {
		return user.NewResetPasswordBadRequest().WithPayload(&user.ResetPasswordBadRequestBody{
			Message: "null user",
		})
	}

	if u.Password != body.Password {
		return user.NewResetPasswordBadRequest().WithPayload(&user.ResetPasswordBadRequestBody{
			Message: "wrong password",
		})
	}
	// u.Temp = "111111"
	if u.Temp != body.VerificationCode {
		return user.NewResetPasswordBadRequest().WithPayload(&user.ResetPasswordBadRequestBody{
			Message: "wrong verification code",
		})
	}

	u.Password = body.NewPassword
	u.PwReset = true
	err := db.Save(&u).Error
	if err != nil {
		return user.NewResetPasswordBadRequest().WithPayload(&user.ResetPasswordBadRequestBody{
			Message: err.Error(),
		})
	}
	return user.NewResetPasswordOK().WithPayload(&user.ResetPasswordOKBody{
		Message: "success",
	})
}

// RegisterUserHandler is API handler for /users/register POST
func RegisterUserHandler(params user.RegisterUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()
	var queryUser models.User
	db.Where("username = ?", body.Username).First(&queryUser)
	if queryUser.Username != "" {
		return user.NewRegisterUserBadRequest().WithPayload(&user.RegisterUserBadRequestBody{Message: "Username already used!"})
	}

	newUser := models.User{
		City:        body.City,
		CompanyName: body.CompanyName,
		Country:     body.Country,
		Email:       body.Email,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Password:    body.Password,
		Phone:       body.Phone,
		Position:    body.Position,
		Priority:    models.UserPriorityLow,
		Username:    body.Username,
		PwReset:     false,
	}

	err := db.Create(&newUser).Error
	if err != nil {
		return user.NewRegisterUserBadRequest()
	}

	res := &user.RegisterUserOKBodyUserInfo{
		City:        body.City,
		CompanyName: body.CompanyName,
		Country:     body.Country,
		Email:       body.Email,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Password:    body.Password,
		Phone:       body.Phone,
		Position:    body.Position,
		Priority:    models.UserPriorityLow,
		Username:    body.Username,
	}

	return user.NewRegisterUserOK().WithPayload(&user.RegisterUserOKBody{UserInfo: res})
}

// UserLoginHandler is API handler for /users/login POST
func UserLoginHandler(params user.UserLoginParams) middleware.Responder {
	body := params.ReqBody

	db := config.GetDB()
	var queryUser models.User
	db.Where("Username = ?", body.Username).First(&queryUser)
	if queryUser.Username == "" {
		return user.NewUserLoginUnauthorized()
	} else if queryUser.Password != body.Password {
		return user.NewUserLoginUnauthorized()
	}

	expirationTime := time.Now().Add(expireTime)
	claims := Claims{
		ID: queryUser.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := config.GetConfig().SecretKey
	signedToken, _ := token.SignedString([]byte(secretKey))

	res := user.UserLoginOKBodyUserInfo{Priority: queryUser.Priority, Username: queryUser.Username,
		OrgName: queryUser.OrgName, PwReset: fmt.Sprintf("%t", queryUser.PwReset), Role: queryUser.Role}

	return user.NewUserLoginOK().WithPayload(&user.UserLoginOKBody{Token: signedToken, UserInfo: &res})
}

// GetUserProfileHandler is API handler for /users/profile GET
func GetUserProfileHandler(params user.GetUserProfileParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewGetUserProfileUnauthorized()
	}

	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var u models.User
	if db.Where("id = ?", uid).First(&u).RowsAffected == 0 {
		return user.NewGetUserProfileNotFound()
	}

	res := user.GetUserProfileOKBodyResults{
		City:        u.City,
		CompanyName: u.CompanyName,
		Country:     u.Country,
		Email:       u.Email,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Phone:       u.Phone,
		Position:    u.Position,
		Priority:    u.Priority,
		Username:    u.Username,
	}

	return user.NewGetUserProfileOK().WithPayload(&user.GetUserProfileOKBody{
		Message: "success",
		Results: &res,
	})
}

// UpdateUserProfileHandler is API handler for /users/profile PUT
func UpdateUserProfileHandler(params user.UpdateUserProfileParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewUpdateUserProfileUnauthorized()
	}

	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	db := config.GetDB()
	var u models.User
	if db.Where("id = ?", uid).First(&u).RowsAffected == 0 {
		return user.NewUpdateUserProfileNotFound()
	}

	u.City = body.City
	u.CompanyName = body.CompanyName
	u.Country = body.Country
	u.Email = body.Email
	u.FirstName = body.FirstName
	u.LastName = body.LastName
	u.Phone = body.Phone
	u.Position = body.Position

	err := db.Save(&u).Error
	if err != nil {
		return user.NewUpdateUserProfileNotFound().WithPayload(&user.UpdateUserProfileNotFoundBody{
			Message: err.Error(),
		})
	}

	return user.NewUpdateUserProfileOK().WithPayload(&user.UpdateUserProfileOKBody{
		Message: "success",
	})
}

func AddUserHandler(params user.AddUserParams) middleware.Responder {
	godotenv.Load("/.env")
	OFFICIAL_EMAIL := os.Getenv("OFFICIAL_EMAIL")
	OFFICIAL_PASSWORD := os.Getenv("OFFICIAL_PASSWORD")
	if !VerifyUser(params.HTTPRequest) {
		return user.NewAddUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	db := config.GetDB()

	if body.Role != "SITE_ADMIN" && body.Role != "ORG_ADMIN" && body.Role != "USER" {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "User Role Invalid. Could only be SITE_ADMIN/ORG_ADMIN/USER",
		})
	}
	var orgNew models.Org
	var userOld models.User
	if db.Where("username = ?", body.Name).First(&userOld).RowsAffected == 1 {
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "User Name Invalid.",
		})
	}
	if db.Where("org_name = ?", body.OrgName).First(&orgNew).RowsAffected == 0 {
		// invalid org
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "Org Name Invalid.",
		})
	}

	pw, _ := password.Generate(10, 4, 0, false, false)
	code, _ := password.Generate(6, 6, 0, false, false)
	fmt.Println("password generated!!!")
	newUser := models.User{
		Username: body.Name,
		Role:     body.Role,
		Email:    body.Email,
		PwReset:  false,
		Phone:    body.Phone,
		OrgName:  body.OrgName,
		Password: pw,
		Temp:     code,
	}
	if db.Unscoped().Where("username = ?", body.Name).First(&userOld).RowsAffected == 1 {
		//delete already deleted user info permently when new user created
		db.Unscoped().Delete(&userOld)
	}

	if err := db.Create(&newUser).Error; err != nil {
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "Insert row DB error: " + err.Error(),
		})
	}

	newLog := models.Log{
		UserID:    uid,
		Operation: "Add User:" + body.Name,
		Time:      time.Now(),
		Status:    "Succeed",
	}
	fmt.Println("test add user!!!")
	fmt.Println("start send!!!")
	m := gomail.NewMessage()
	m.SetHeader("From", OFFICIAL_EMAIL)
	m.SetHeader("To", body.Email)
	m.SetHeader("Subject", "CloudTides Default Password")
	m.SetBody("text/plain", "CloudTides has registered an account for you. Your username is: " + body.Name + "\nYour login password for CloudTides is: "+pw+"\nPlease login to CloudTides platform and reset the password. Your verification code is: "+code)
	fmt.Println("m set!!!")
	// d := gomail.NewDialer("smtp.gmail.com", 587, OfficialEmail, OfficialPassword)
	d := gomail.NewDialer("smtp.qiye.aliyun.com", 25, OFFICIAL_EMAIL, OFFICIAL_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	i := 0
	for ; i < 5; i++ {
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	if i == 5 {
		fmt.Println("fail")
		db.Delete(&newUser)
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "fail!!!",
		})
	}
	// if err := d.DialAndSend(m); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	fmt.Println("success!!!")

	if err := db.Create(&newLog).Error; err != nil {
		return user.NewAddUserForbidden().WithPayload(&user.AddUserForbiddenBody{
			Message: "Insert Log DB error: " + err.Error(),
		})
	}
	return user.NewAddUserOK().WithPayload(&user.AddUserOKBody{
		Message: "succeed",
	})
}

func ListUserHandler(params user.ListUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewListUserUnauthorized()
	}
	var users []*models.User
	db := config.GetDB()
	db.Order("id").Find(&users)
	var response []*user.ListUserOKBodyItems0
	for _, tmpUser := range users {
		newResult := user.ListUserOKBodyItems0{
			Email:   tmpUser.Email,
			ID:      int64(tmpUser.ID),
			Name:    tmpUser.Username,
			Phone:   tmpUser.Phone,
			Role:    tmpUser.Role,
			OrgName: tmpUser.OrgName,
		}

		response = append(response, &newResult)
	}
	return user.NewListUserOK().WithPayload(response)
}

func ListUserOfOrgHandler(params user.ListUserOfOrgParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewListUserUnauthorized()
	}
	var users []*models.User
	db := config.GetDB()
	db.Where("org_name = ?", params.OrgName).Order("id").Find(&users)
	var response []*user.ListUserOKBodyItems0
	for _, tmpUser := range users {
		newResult := user.ListUserOKBodyItems0{
			Email:   tmpUser.Email,
			ID:      int64(tmpUser.ID),
			Name:    tmpUser.Username,
			Phone:   tmpUser.Phone,
			Role:    tmpUser.Role,
			OrgName: tmpUser.OrgName,
		}

		response = append(response, &newResult)
	}
	return user.NewListUserOK().WithPayload(response)
}

func ModifyUserHandler(params user.ModifyUserParams) middleware.Responder {
	OFFICIAL_EMAIL := os.Getenv("OFFICIAL_EMAIL")
	OFFICIAL_PASSWORD := os.Getenv("OFFICIAL_PASSWORD")
	if !VerifyUser(params.HTTPRequest) {
		return user.NewModifyUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	var pol models.User
	db := config.GetDB()
	if db.Where("id = ?", params.ID).First(&pol).RowsAffected == 0 {
		return user.NewModifyUserNotFound().WithPayload(&user.ModifyUserNotFoundBody{
			Message: "user with id " + strconv.FormatInt(params.ID, 10) + " is not found in database",
		})
	}

	//pol.Org.OrgName = body.Org
	if body.Role != "SITE_ADMIN" && body.Role != "ORG_ADMIN" && body.Role != "USER" {
		return user.NewModifyUserForbidden().WithPayload(&user.ModifyUserForbiddenBody{
			Message: "User Role Invalid. Could only be SITE_ADMIN/ORG_ADMIN/USER",
		})
	}
	pol.Username = body.Name
	pol.Phone = body.Phone
	pol.Email = body.Email
	pol.Role = body.Role
	err := db.Save(&pol).Error
	if err != nil {
		return user.NewModifyUserForbidden().WithPayload(&user.ModifyUserForbiddenBody{
			Message: "Insert row DB error: " + err.Error(),
		})
	}

	newLog := models.Log{
		UserID:    uid,
		Operation: "Modify User:" + body.Name,
		Time:      time.Now(),
		Status:    "Succeed",
	}
	if err = db.Create(&newLog).Error; err != nil {
		return user.NewModifyUserForbidden().WithPayload(&user.ModifyUserForbiddenBody{
			Message: "Insert Log DB error: " + err.Error(),
		})
	}

	fmt.Println("test update user!!!")
	fmt.Println("start send!!!")
	m := gomail.NewMessage()
	m.SetHeader("From", OFFICIAL_EMAIL)
	m.SetHeader("To", body.Email)
	m.SetHeader("Subject", "CloudTides User Information Updated")
	m.SetBody("text/plain", "Your Cloudtides account has been updated. Your username is "+body.Name+" and your email is "+body.Email+".")
	fmt.Println("m set!!!")
	// d := gomail.NewDialer("smtp.gmail.com", 587, OFFICIAL_EMAIL, OFFICIAL_PASSWORD)
	d := gomail.NewDialer("smtp.qiye.aliyun.com", 25, OFFICIAL_EMAIL, OFFICIAL_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// if err := d.DialAndSend(m); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	i := 0
	for ; i < 5; i++ {
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	if i == 5 {
		fmt.Println("fail")
		return user.NewSendVerificationBadRequest().WithPayload(&user.SendVerificationBadRequestBody{
			Message: "fail!!!",
		})
	}
	fmt.Println("success!!!")

	return user.NewModifyUserOK().WithPayload(&user.ModifyUserOKBody{
		Message: "success",
	})

}

func DeleteUserHandler(params user.DeleteUserParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return user.NewDeleteUserUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var pol models.User
	if db.Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return user.NewDeleteUserNotFound().WithPayload(&user.DeleteUserNotFoundBody{
			Message: "user with id " + strconv.FormatInt(params.ID, 10) + " is not found in database",
		})
	}

	newLog := models.Log{
		UserID:    uid,
		Operation: "Delete User with Id: " + strconv.FormatInt(params.ID, 10),
		Time:      time.Now(),
		Status:    "Succeed",
	}

	if err := db.Create(&newLog).Error; err != nil {
		return user.NewDeleteUserForbidden().WithPayload(&user.DeleteUserForbiddenBody{
			Message: "delete user failed: " + err.Error(),
		})
	}

	return user.NewDeleteUserOK().WithPayload(&user.DeleteUserOKBody{
		Message: "success",
	})
}
