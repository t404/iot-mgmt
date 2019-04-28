package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"iot-mgmt/services"
)

type UserController struct {
}

// @Summary 获取当前用户信息
// @Accept json
// @Tags 用户
// @Security Bearer
// @Produce  json
// @Param id path string true "查询用户"
// @Router /user/{id} [get]
// @Success 200 {object} models.User
func (t *UserController) Find(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "find user",
	})
}

// @Summary 创建
// @Accept json
// @Tags 用户
// @Produce json
// @Param name body services.CreateUserRequest true "创建数据"
// @Resource nil
// @Router /users [post]
// @Success 200 {object} models.User
func (t *UserController) Create(c *gin.Context) {
	userService := services.NewCompanyService()
	var request services.CreateUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if model, err := userService.Create(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, model)
	}
}

// @Summary 更新
// @Accept json
// @Tags 用户
// @Produce  json
// @Param id path string true "更新数据"
// @Resource nil
// @Router /user/{id} [put]
// @Success 201 {object} models.User
func (t *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "update user",
	})
}

// @Summary 查询用户列表
// @Accept json
// @Tags 用户
// @Produce  json
// @Resource nil
// @Router /users [get]
// @Success 201 {object} models.User
func (t *UserController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page":       1,
		"total_page": 2,
		"total":      30,
		"per_page":   15,
		"list": []map[string]interface{}{
			{
				"id":   1,
				"name": "list user",
			},
		},
	})
}

// @Summary 删除
// @Accept json
// @Tags 用户
// @Produce  json
// @Param id body services.DeleteUserRequest true "删除数据"
// @Resource nil
// @Router /users [delete]
// @Success 201 {object} models.User
func (t *UserController) Delete(c *gin.Context) {
	userService := services.NewCompanyService()
	var request services.DeleteUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if model, err := userService.Delete(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, model)
	}
}
