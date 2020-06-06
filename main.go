package main

import (
	"gql-metrics/structs"
	"gql-metrics/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/language/parser"
)

// A GQLRequest struct is a single GraphQL server request
type GQLRequest struct {
	Query string `json:"query"`
}

var operations = make([]structs.Operation, 0)

func getOperations(c *gin.Context) {
	c.JSON(http.StatusOK, operations)
}

func addQuery(c *gin.Context) {
	var gqlRequest GQLRequest
	err := c.BindJSON(&gqlRequest)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query is not a valid GrapQL query",
		})

		return
	}

	parseParams := parser.ParseParams{
		Source:  gqlRequest.Query,
		Options: parser.ParseOptions{NoLocation: true, NoSource: true},
	}

	AST, err := parser.Parse(parseParams)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse query",
		})

		return
	}

	operationDefinition, err := utils.GetOperationDefinitionFromDocument(AST)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	var fields []structs.Field

	utils.GetFieldsFromOperationDefinitionSelectionSet(operationDefinition.SelectionSet, &fields)

	operation := structs.Operation{
		Type:   operationDefinition.Operation,
		Fields: fields,
	}

	operations = append(operations, operation)

	c.JSON(http.StatusCreated, gin.H{
		"message": "GraphQL query registered successfully",
	})
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	api := router.Group("/api")
	{
		api.GET("/operations", getOperations)
		api.POST("/queries", addQuery)
	}

	router.Run(":5000")
}
