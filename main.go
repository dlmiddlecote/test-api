package main

import "github.com/gin-gonic/gin"
import "errors"
import "strconv"

type resource struct {
	ID       int
	Forename string
	Surname  string
}

var resources map[int]resource

var resourceNotFound = errors.New("Resource Not Found")

func addResource(forename, surname string) int {
	id := len(resources) + 1
	resources[id] = resource{
		id,
		forename,
		surname,
	}
	return id
}

func getResource(id int) (resource, error) {
	r, ok := resources[id]
	if !ok {
		return resource{}, resourceNotFound
	}
	return r, nil
}

func getResourcesByFilter(match func(resource) bool) []resource {
	rs := []resource{}
	for _, r := range resources {
		if match(r) {
			rs = append(rs, r)
		}
	}
	return rs
}

func createJSONRep(r resource) gin.H {
	return gin.H{
		"id":       r.ID,
		"forename": r.Forename,
		"surname":  r.Surname,
	}
}

func init() {
	resources = make(map[int]resource)

	// Populate test resources
	rs := [][]string{
		[]string{"daniel", "middlecote"},
		[]string{"paul", "middlecote"},
		[]string{"federico", "figus"},
		[]string{"toby", "dunn"},
		[]string{"gisela", "rossi"},
	}

	for _, r := range rs {
		addResource(r[0], r[1])
	}
}

func main() {
	r := gin.Default()

	r.GET("/resources", func(c *gin.Context) {
		rs := []gin.H{}

		forename := c.Query("forename")
		surname := c.Query("surname")

		var match func(resource) bool

		if forename != "" && surname != "" {
			// Both forename and surname in query
			match = func(r resource) bool {
				return r.Forename == forename && r.Surname == surname
			}
		} else if forename != "" {
			// Just forename in query
			match = func(r resource) bool {
				return r.Forename == forename
			}
		} else if surname != "" {
			// Just surname in query
			match = func(r resource) bool {
				return r.Surname == surname
			}
		} else {
			// No query
			match = func(_ resource) bool {
				return true
			}
		}

		for _, r := range getResourcesByFilter(match) {
			rs = append(rs, createJSONRep(r))
		}

		c.JSON(200, rs)
	})

	r.GET("/resources/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(400, "id must be an integer")
			return
		}

		r, err := getResource(id)

		if err != nil {
			c.String(404, err.Error())
			return
		}

		c.JSON(200, createJSONRep(r))
	})

	r.POST("/resources", func(c *gin.Context) {
		var r resource
		c.BindJSON(&r)
		id := addResource(r.Forename, r.Surname)
		c.JSON(201, createJSONRep(resources[id]))
	})

	r.Run(":1234")
}
