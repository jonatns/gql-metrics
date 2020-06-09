package utils

import (
	"errors"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/jonatns/gql-metrics/structs"
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
					var arguments []structs.Argument

					for _, argument := range selection.Arguments {
						arguments = append(arguments, structs.Argument{Name: argument.Name.Value, Kind: argument.Value.GetKind()})
					}

					*fields = append(*fields, structs.Field{Name: selection.Name.Value, Fields: newFields, Arguments: arguments})
				}
			}
		}
	}
}
