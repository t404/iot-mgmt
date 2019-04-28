package controllers

import "github.com/gin-gonic/gin"

// 列表接口
type ControllerList interface {
    List(c *gin.Context)
}
// 创建接口
type ControllerCreate interface {
    Create(c *gin.Context)
}
// 更新接口
type ControllerUpdate interface {
    Update(c *gin.Context)
}
// 查找单条记录接口
type ControllerFind interface {
    Find(c *gin.Context)
}
// 删除接口
type ControllerDelete interface {
    Delete(c *gin.Context)
}
// 控制器接口
type Controller interface {
    ControllerList
    ControllerCreate
    ControllerUpdate
    ControllerFind
    ControllerDelete
}