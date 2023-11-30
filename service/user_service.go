package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// Create User
//
// @Summary	Create User
// @Tags		User Service
// @Param		name		query		string	true	"User Name"
// @Param		password	query		string	true	"Password"
// @Param		repassword	query		string	true	"Password Again"
// @Param		phone		query		string	true	"User Phone Number"
// @Param		email		query		string	true	"User Email Address"
// @Success	200			{string}	json{"code", "message"}
// @Router		/user/createUser [get]
func CreateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Query("name")
	user.Phone = ctx.Query("phone")
	user.Email = ctx.Query("email")
	password, rePassword := ctx.Query("password"), ctx.Query("repassword")
	if password != rePassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "password not the same",
		})
		return
	}
	// user.Password = password
	if len(models.GetUserByName(user.Name)) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "name already exit",
		})
		return
	} else if len(models.GetUserByPhone(user.Phone)) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "phone already exit",
		})
		return
	} else if len(models.GetUserByEmail(user.Email)) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "email already exit",
		})
		return
	}

	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(password, salt)

	models.CreateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": user,
	})
}

// Delete User
//
// @Summary	Delete User
// @Tags		User Service
// @Param		id	query		string	true	"User ID"
// @Success	200	{string}	json{"code", "message"}
// @Router		/user/deleteUser [get]
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": user,
	})
}

// Update User
//
// @Summary	Update User
// @Tags		User Service
// @Param		id		query		string	true	"User ID"
// @Param		name	query		string	true	"User Name"
// @Param		phone	query		string	true	"User Phone Number"
// @Param		email	query		string	true	"User Email Address"
// @Success	200		{string}	json{"code", "message"}
// @Router		/user/updateUser [get]
func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	user.ID = uint(id)
	user.Name = ctx.Query("name")
	user.Phone = ctx.Query("phone")
	user.Email = ctx.Query("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": fmt.Sprintln(err),
		})
		return
	}
	models.UpdateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": user,
	})
}

//	 Get login User data
//		@Tags		User Service
//		@Param		name		query		string	true	"User Name"
//		@Param		password	query		string	true	"Password"
//		@Success	200			{string}	json{"code", "message"}
//		@Router		/user/loginVal [get]
func LoginValidate(ctx *gin.Context) {
	name, tPwd := ctx.Query("name"), ctx.Query("password")

	data := models.GetUserByName(name)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Name is not existing",
		})
		return
	}
	user := data[0]
	if !utils.ValidatePassword(tPwd, user.Salt, user.Password) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Password is not correct",
		})
		return
	}

	user.Identity = fmt.Sprintf("%d", time.Now().Unix())
	models.UpdateUserTocken(user)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "login sucesss",
	})
}
