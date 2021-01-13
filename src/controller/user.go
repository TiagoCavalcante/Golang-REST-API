package controller

import (
	"encoding/json"
	"strconv"

	"config"
	"entities"
	"models"

	"github.com/valyala/fasthttp"
)

var database, _ = config.GetDatabase()	
var userModel = models.UserModel{database}

func GetUsers(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	users, err := userModel.FindAllUsers()

	if err != nil {
		context.SetStatusCode(500)

		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		
		context.WriteString(string(response))

		return
	}

	response, _ := json.Marshal(users)
		
	context.WriteString(string(response))
}

func GetUser(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)
	user, err := userModel.FindUserById(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	emptyUser := entities.User{}

	if user == emptyUser {
		context.SetStatusCode(404)

		return
	}
		
	response, _ := json.Marshal(user)

	context.WriteString(string(response))
}

func CreateUser(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	user := entities.User{}

	err := json.Unmarshal(context.PostBody(), &user)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.SetStatusCode(400)
		context.WriteString(string(response))

		return
	}
		
	err = userModel.CreateUser(&user)

	if err != nil {
		context.SetStatusCode(500)

		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		
		context.WriteString(string(response))

		return
	}
		
	response, _ := json.Marshal(map[string]int64{"id": user.Id})

	context.SetStatusCode(201)
	context.WriteString(string(response))
}

func UpdateUser(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")

	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)
	user := entities.User{Id: id}

	err := json.Unmarshal(context.PostBody(), &user)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		
		context.SetStatusCode(400)
		context.WriteString(string(response))

		return
	}
	
	rows, err := userModel.UpdateUser(&user)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
			
		context.WriteString(string(response))

		return
	}
	
	if rows == 0 {
		context.SetStatusCode(404)

		return
	}

	context.SetStatusCode(204)
}

func DeleteUser(context *fasthttp.RequestCtx) {
	context.SetContentType("application/json")
	
	idString, _ := context.UserValue("id").(string)
	id, _ := strconv.ParseInt(idString, 10, 64)

	user, err := userModel.FindUserById(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})

		context.WriteString(string(response))

		return
	}

	emptyUser := entities.User{}

	if user == emptyUser {
		context.SetStatusCode(404)

		return
	}

	_, err = userModel.DeleteUser(id)

	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		
		context.WriteString(string(response))

		return
	}

	context.SetStatusCode(204)
}