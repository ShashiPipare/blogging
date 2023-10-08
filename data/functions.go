package data

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func New(c *fiber.Ctx) (a *model) {
	a = &model{
		c,
	}
	return
}

func Convert(fromAddress any, toAddress any) (err error) {
	bytes, err := json.Marshal(fromAddress)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, toAddress)
	if err != nil {
		return
	}
	return
}

func ParseArray(from Object, key string) (result []any, err error) {
	switch data := from[key].(type) {
	case []any:
		result = data
	case []map[string]any:
		for _, element := range data {
			result = append(result, element)
		}
	case []Object:
		for _, element := range data {
			result = append(result, element)
		}
	case primitive.A:
		result = data
	case primitive.M:
		for _, element := range data {
			result = append(result, element)
		}
	case []string:
		for _, element := range data {
			result = append(result, element)
		}
	default:
		err = ErrParseArray
	}
	return
}

func ParseAsObjectArray(from Object, key string) (result []Object, err error) {
	switch data := from[key].(type) {
	case []map[string]any:
		for _, element := range data {
			result = append(result, element)
		}
	case []Object:
		result = data
	default:
		err = Convert(data, &result)
	}
	return
}
