package blog

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/data"
)

func create(c *fiber.Ctx) (err error) {
	a := data.New(c)
	payload, err := a.Payload()
	if err != nil {
		a.Error(err)
	}
	t := time.Now().UTC()
	userIDHex, ok := payload["user_id"].(string)
	if !ok {
		return a.Error(ErrInvalidUserID)
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return a.Error(err)
	}
	body, ok := payload["body"].(string)
	if !ok {
		return a.Error(ErrInvalidString)
	}
	array, err := data.ParseArray(payload, "users")
	if err != nil {
		return a.Error(err)
	}
	if len(array) == 0 {
		return a.Error(ErrEmptyArray)
	}
	users := []User{}
	for i := range array {
		user := User{}
		err = data.Convert(array[i], &user)
		if err != nil {
			err = fmt.Errorf("error in converting users array at index:%d : %v", i, err)
			return a.Error(err)
		}
		users = append(users, user)
	}
	blog := &Blog{}
	blog.Body = body
	blog.Users = users
	blog.Likes = []Like{}
	blog.Comments = []Comment{}
	blog.Type = "parent"
	err = blog.create(userID, t)
	if err != nil {
		return a.Error(err)
	}
	return a.Data(blog)
}
func update(c *fiber.Ctx) (err error) {
	a := data.New(c)
	payload, err := a.Payload()
	if err != nil {
		a.Error(err)
	}
	t := time.Now().UTC()
	blogIDHex, ok := payload["blog_id"].(string)
	if !ok {
		return a.Error(ErrInvalidString)
	}
	blogID, err := primitive.ObjectIDFromHex(blogIDHex)
	if err != nil {
		return a.Error(err)
	}
	userIDHex, ok := payload["user_id"].(string)
	if !ok {
		return a.Error(ErrInvalidUserID)
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return a.Error(err)
	}
	oldBlog := &Blog{}
	err = oldBlog.fetchBlog(userID, blogID)
	if err != nil {
		return a.Error(err)
	}
	AllUsers := make(map[primitive.ObjectID]User)
	for i := range oldBlog.Users {
		AllUsers[oldBlog.Users[i].UserID] = oldBlog.Users[i]
	}
	body, ok := payload["body"].(string)
	if !ok {
		return a.Error(ErrInvalidString)
	}
	array, err := data.ParseArray(payload, "new_users")
	if err != nil {
		return a.Error(err)
	}
	newUsers := []User{}
	for i := range array {
		user := User{}
		err = data.Convert(array[i], &user)
		if err != nil {
			err = fmt.Errorf("error in converting users array at index:%d : %v", i, err)
			return a.Error(err)
		}
		_, ok := AllUsers[user.UserID]
		if !ok {
			newUsers = append(newUsers, user)
			continue
		}
		//check for deleted and update existing user
	}
	blog := &Blog{}
	err = blog.update(userID, blogID, body, newUsers, t)
	if err != nil {
		return a.Error(err)
	}
	return a.Data(blog)
}
func delete(c *fiber.Ctx) (err error) {
	a := data.New(c)
	payload, err := a.Payload()
	if err != nil {
		a.Error(err)
	}
	userIDHex, ok := payload["user_id"].(string)
	if !ok {
		return a.Error(ErrInvalidUserID)
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return a.Error(err)
	}
	blogIDHex, ok := payload["blog_id"].(string)
	if !ok {
		return a.Error(ErrInvalidString)
	}
	blogID, err := primitive.ObjectIDFromHex(blogIDHex)
	if err != nil {
		return a.Error(err)
	}
	err = deleteBlog(userID, blogID)
	if err != nil {
		return a.Error(err)
	}
	return a.True()
}
func getOne(c *fiber.Ctx) (err error) {
	a := data.New(c)
	payload, err := a.Payload()
	if err != nil {
		a.Error(err)
	}
	userIDHex, ok := payload["user_id"].(string)
	if !ok {
		return a.Error(ErrInvalidUserID)
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return a.Error(err)
	}
	blogIDHex, ok := payload["blog_id"].(string)
	if !ok {
		return a.Error(ErrInvalidString)
	}
	blogID, err := primitive.ObjectIDFromHex(blogIDHex)
	if err != nil {
		return a.Error(err)
	}
	blog := &Blog{}
	err = blog.fetchBlog(userID, blogID)
	if err != nil {
		return a.Error(err)
	}
	return a.Data(blog)
}
func getAll(c *fiber.Ctx) (err error) {
	a := data.New(c)
	payload, err := a.Payload()
	if err != nil {
		a.Error(err)
	}
	userIDHex, ok := payload["user_id"].(string)
	if !ok {
		return a.Error(ErrInvalidUserID)
	}
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return a.Error(err)
	}
	// array, err := data.ParseArray(payload, "blog_ids")
	// if err != nil {
	// 	return a.Error(err)
	// }
	// IDs := []primitive.ObjectID{}
	// for i := range array {
	// 	blogIDHex, ok := array[i].(string)
	// 	if !ok {
	// 		return a.Error(ErrInvalidString)
	// 	}
	// 	blogID, err := primitive.ObjectIDFromHex(blogIDHex)
	// 	if err != nil {
	// 		return a.Error(err)
	// 	}
	// 	IDs = append(IDs, blogID)
	// }
	blogs, err := fetchBlogs(userID)
	if err != nil {
		return a.Error(err)
	}
	return a.Data(blogs)
}
