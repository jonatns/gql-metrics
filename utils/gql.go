package utils

import (
	"errors"
	"gql-metrics/structs"
	"log"

	"github.com/graphql-go/graphql/language/ast"
)

// GetOperationDefinitionFromDocument returns the operation definition of a GraphQL Document
func GetOperationDefinitionFromDocument(document *ast.Document) (*ast.OperationDefinition, error) {
	for _, definition := range document.Definitions {
		switch definition := definition.(type) {
		case *ast.OperationDefinition:
			return definition, nil
		default:
			return nil, errors.New("Invalid kind: " + definition.GetKind())
		}
	}

	return nil, errors.New("Invalid GraphQL document definitions")
}

// GetFieldsFromOperationDefinitionSelectionSet recursively retrieves the fields of a selectionSet
func GetFieldsFromOperationDefinitionSelectionSet(selectionSet *ast.SelectionSet, fields *[]structs.Field) {
	if selectionSet != nil {
		for _, selection := range selectionSet.Selections {
			switch selection := selection.(type) {
			case *ast.Field:
				var newFields []structs.Field
				GetFieldsFromOperationDefinitionSelectionSet(selection.SelectionSet, &newFields)
				if selection.Name.Value != "query" && selection.Name.Value != "mutation" {
					log.Println(selection.)
					*fields = append(*fields, structs.Field{Name: selection.Name.Value, Fields: newFields})
				}
			}
		}
	}
}
