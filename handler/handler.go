package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CommandHandler(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("[Command]cannot parse body command")
	}

	var requestData RequestData
	err = json.Unmarshal(jsonData, &requestData)
	if err != nil {
		fmt.Println("[Command]cannot parse json command")
	}

	fmt.Println(fmt.Sprintf("%+v", requestData))
	c.JSON(http.StatusOK, "ok")
}
