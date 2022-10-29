package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	apps_repository "vessl/modules/apps-repository"
	mqttcloudconfig "vessl/modules/apps/mqtt-cloud-connector/config"
	opcuaconfig "vessl/modules/apps/opcua-mqtt-connector/config"
	"vessl/modules/configurator"
	"vessl/modules/containers"
	"vessl/modules/images"
	"vessl/modules/networks"
	"vessl/modules/system"
	"vessl/modules/users"
	"vessl/modules/volumes"

	"github.com/gin-gonic/gin"
)

// ///////////CG-APPS CONFIGURATION HANDLERS////////////////////
func GetConfigHandler(c *gin.Context) {
	appName := c.Param("appName")
	switch appName {
	case "mqtt-cloud-connector":
		c.JSON(http.StatusOK, configurator.GetMccConfig())
		return
	case "opcua-mqtt-connector":
		c.JSON(http.StatusOK, configurator.GetOpcuaConfig())
		return
	}
	c.JSON(http.StatusBadRequest, "App not found")
}

func SetConfigHandler(c *gin.Context) {
	appName := c.Param("appName")
	switch appName {
	case "mqtt-cloud-connector":
		configFile, statusCode, err := convertHTTPBodyMccConfig(c.Request.Body)
		if err != nil {
			c.JSON(statusCode, err)
			return
		}
		c.JSON(statusCode, configurator.SetMccConfig(configFile))
		return
	case "opcua-mqtt-connector":
		configFile, statusCode, err := convertHTTPBodyOpcuaConfig(c.Request.Body)
		if err != nil {
			c.JSON(statusCode, err)
			return
		}
		c.JSON(statusCode, configurator.SetOpcuaConfig(configFile))
		return
	}
	c.JSON(http.StatusBadRequest, "App not found")
}

// ///////////CONTAINERS HANDLERS////////////////////
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

func GetLogsHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, containers.Logs(Id))
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

// ///////////VOLUMES HANDLERS////////////////////
func GetVolumesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, volumes.GetVolumes())
}

func RemoveVolumeHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, volumes.RemoveVolume(Id))
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

// ////////////USERS HANDLERS/////////////////////
func GetUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users.GetUsers())
}

func UpdateUserHandler(c *gin.Context) {
	User, statusCode, err := convertHTTPBodyUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, users.UpdateUser(User))
}

func AddUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users.AddUser())
}

func DeleteUserHandler(c *gin.Context) {
	Id := c.Param("Id")
	c.JSON(http.StatusOK, users.DeleteUser(Id))
}

func ValidateUserHandler(c *gin.Context) {
	User, statusCode, err := convertHTTPBodyUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, users.Validate(User))
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

func convertHTTPBodyMccConfig(httpBody io.ReadCloser) (mqttcloudconfig.Config, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return mqttcloudconfig.Config{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var Config mqttcloudconfig.Config
	err = json.Unmarshal(body, &Config)
	if err != nil {
		return mqttcloudconfig.Config{}, http.StatusBadRequest, err
	}
	return Config, http.StatusOK, nil
}

func convertHTTPBodyOpcuaConfig(httpBody io.ReadCloser) (opcuaconfig.Config, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return opcuaconfig.Config{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var Config opcuaconfig.Config
	err = json.Unmarshal(body, &Config)
	if err != nil {
		return opcuaconfig.Config{}, http.StatusBadRequest, err
	}
	return Config, http.StatusOK, nil
}

func convertHTTPBodyAppTemplate(httpBody io.ReadCloser) (apps_repository.Template, int, error) {
	body, err := ioutil.ReadAll(httpBody)
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
	body, err := ioutil.ReadAll(httpBody)
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

func convertHTTPBodyUsers(httpBody io.ReadCloser) (users.Users, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return users.Users{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var Users users.Users
	err = json.Unmarshal(body, &Users)
	if err != nil {
		return users.Users{}, http.StatusBadRequest, err
	}
	return Users, http.StatusOK, nil
}

func convertHTTPBodyUser(httpBody io.ReadCloser) (users.User, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return users.User{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	var User users.User
	err = json.Unmarshal(body, &User)
	if err != nil {
		return users.User{}, http.StatusBadRequest, err
	}
	return User, http.StatusOK, nil
}
