package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	apps_repository "vessl/modules/apps-repository"
	"vessl/modules/containers"
	db "vessl/modules/database"
	"vessl/modules/images"
	"vessl/modules/networks"
	"vessl/modules/system"
	"vessl/modules/volumes"

	"github.com/gin-gonic/gin"
)

// ////////////CONTAINERS HANDLERS////////////////////
func GetContainersHandler(c *gin.Context) {
	networkName := c.Param("networkName")
	c.JSON(http.StatusOK, containers.GetContainers(networkName))
}

func GetAppRepositoryHandler(c *gin.Context) {
	UserId := c.Param("UserId")
	c.JSON(http.StatusOK, apps_repository.GetApps(UserId))
}

func InstallContainerHandler(c *gin.Context) {
	AppTemplate, statusCode, err := convertHTTPBodyAppTemplate(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, containers.InstallContainer(AppTemplate))
}

func StartContainerHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.StartContainer(Id))
}

func StopContainerHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.StopContainer(Id))
}

func RestartContainerHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.RestartContainer(Id))
}

func RemoveContainerHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.RemoveContainer(Id))
}

func InspectContainerHandler(c *gin.Context) {
	Id := c.Param("Id")
	ci, err := containers.InspectContainer(Id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, ci)
}

func GetContainerStatsHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.GetContainerStats(Id))
}

func GetLogsHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.Logs(Id))
}

func GetCompleteStatsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, containers.GetCompleteStats())
}

func GetDockerServerInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, containers.GetDockerServerInfo())
}

// ///////////IMAGES HANDLERS////////////////////
func GetImagesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, images.GetImages())
}

func RemoveImageHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, images.RemoveImage(Id))
}

func InspectImageHandler(c *gin.Context) {
	Id := c.Param("Id")
	ii, err := images.InspectImage(Id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, ii)
}

// ///////////VOLUMES HANDLERS////////////////////
func GetVolumesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, volumes.GetVolumes())
}

func RemoveVolumeHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, volumes.RemoveVolume(Id))
}

func InspectVolumeHandler(c *gin.Context) {
	Id := c.Param("Id")
	vi, err := volumes.InspectVolume(Id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, vi)
}

// ///////////NETWORKS HANDLERS////////////////////
func GetNetworksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, networks.GetNetworks())
}

func CreateNetworkHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, networks.CreateNetwork(Id))
}

func RemoveNetworkHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, networks.RemoveNetwork(Id))
}

func InspectNetworkHandler(c *gin.Context) {
	Id := c.Param("Id")
	ni, err := networks.InspectNetwork(Id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, ni)
}

// ////////////USERS HANDLERS/////////////////////
func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetUsers())
}

func UpdateUserHandler(c *gin.Context) {
	User, statusCode, err := convertHTTPBodyUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, db.UpdateUser(User))
}

func AddUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, db.AddUser())
}

func DeleteUserHandler(c *gin.Context) {
	User, statusCode, err := convertHTTPBodyUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, db.DeleteUser(User))
}

func ValidateUserHandler(c *gin.Context) {
	User, statusCode, err := convertHTTPBodyUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, User.Validate())
}

// ////////////SYSTEM HANDLERS////////////////////
func GetNetworkInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.GetNetworkInfo())
}

func SetNetworkInfoHandler(c *gin.Context) {
	InterfaceSet, statusCode, err := convertHTTPBodyInterfaceSet(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, system.SetNetworkInfo(InterfaceSet))
}

func RestartHostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.RestartHost())
}

func ShutDownHostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.ShutDownHost())
}

func GetHostStatsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, system.GetHostStats())
}

///////////////CONVERSIONs OF HTTP BODY TO SPECIFIC STRUCTURES////////////////////////////

func convertHTTPBodyAppTemplate(httpBody io.ReadCloser) (apps_repository.Template, int, error) {
	body, err := io.ReadAll(httpBody)
	if err != nil {
		return apps_repository.Template{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var AppTemplate apps_repository.Template
	err = json.Unmarshal(body, &AppTemplate)
	if err != nil {
		return apps_repository.Template{}, http.StatusBadRequest, err
	}
	return AppTemplate, http.StatusOK, nil
}

func convertHTTPBodyInterfaceSet(httpBody io.ReadCloser) (system.InterfaceSet, int, error) {
	body, err := io.ReadAll(httpBody)
	if err != nil {
		return system.InterfaceSet{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var InterfaceSet system.InterfaceSet
	err = json.Unmarshal(body, &InterfaceSet)
	if err != nil {
		return system.InterfaceSet{}, http.StatusBadRequest, err
	}
	return InterfaceSet, http.StatusOK, nil
}

func convertHTTPBodyUsers(httpBody io.ReadCloser) (db.Users, int, error) {
	body, err := io.ReadAll(httpBody)
	if err != nil {
		return db.Users{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var Users db.Users
	err = json.Unmarshal(body, &Users)
	if err != nil {
		return db.Users{}, http.StatusBadRequest, err
	}
	return Users, http.StatusOK, nil
}

func convertHTTPBodyUser(httpBody io.ReadCloser) (db.User, int, error) {
	body, err := io.ReadAll(httpBody)
	if err != nil {
		return db.User{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var User db.User
	err = json.Unmarshal(body, &User)
	if err != nil {
		return db.User{}, http.StatusBadRequest, err
	}
	return User, http.StatusOK, nil
}
