package internal

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	users          map[int]User = make(map[int]User)
	userIdSeq      int          = 0
	characterIdSeq int          = 0
)

// health check

func GetHealth(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// Users

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, mapToSlice(users))
}

func PostUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Characters = make(map[int]Character)
	userIdSeq++
	u.Id = userIdSeq
	users[userIdSeq] = *u
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if u, found := users[id]; found {
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func PutUser(c echo.Context) error {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if u, found := users[id]; found {
			if err = c.Bind(u); err != nil {
				return err
			}
			c.JSON(http.StatusOK, u)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if _, found := users[id]; found {
			delete(users, id)
			return c.NoContent(http.StatusOK)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

// Characters

func GetCharacters(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userId, err := strconv.Atoi(userIdStr); err == nil {
		if u, found := users[userId]; found {
			return c.JSON(http.StatusOK, mapToSlice(u.Characters))
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func PostCharacter(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userId, err := strconv.Atoi(userIdStr); err == nil {
		if u, found := users[userId]; found {
			char := new(Character)
			if err = c.Bind(char); err != nil {
				return err
			}
			characterIdSeq++
			char.Id = characterIdSeq
			char.UserId = userId
			u.Characters[characterIdSeq] = *char
			return c.JSON(http.StatusCreated, char)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func DeleteCharacter(c echo.Context) error {
	userIdStr := c.Param("userId")
	characterIdStr := c.Param("characterId")
	if userId, err := strconv.Atoi(userIdStr); err == nil {
		if u, found := users[userId]; found {
			if characterId, err := strconv.Atoi(characterIdStr); err == nil {
				if _, found := u.Characters[characterId]; found {
					delete(u.Characters, characterId)
					return c.NoContent(http.StatusOK)
				}
			}
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func mapToSlice(m interface{}) (rtn []interface{}) {
	rtn = make([]interface{}, 0)
	switch m.(type) {
	case map[int]User:
		for id, elem := range m.(map[int]User) {
			elem.Id = id
			rtn = append(rtn, elem)
		}
	case map[int]Character:
		rtn = make([]interface{}, 0)
		for id, elem := range m.(map[int]Character) {
			elem.Id = id
			rtn = append(rtn, elem)
		}
	}
	return
}
