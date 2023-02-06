package resolver

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/iammunir/starwars-graphql-api/helper"
)

type ResolverKey struct {
	Key         string
	selectQuery string
}

func NewResolverKey(key string, selectQuery string) *ResolverKey {
	return &ResolverKey{
		Key:         key,
		selectQuery: selectQuery,
	}
}

func (rk *ResolverKey) String() string {
	return rk.Key
}

func (rk *ResolverKey) Raw() interface{} {
	return rk.Key
}

func (rk *ResolverKey) GetSelectQuery() string {
	return rk.selectQuery
}

type resultLoader struct {
	data interface{}
	err  error
}

func getSelectedFields(params graphql.ResolveParams) ([]string, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) ([]string, error) {
	var selected []string
	for _, s := range selections {
		switch t := s.(type) {
		case *ast.Field:
			selected = append(selected, s.(*ast.Field).Name.Value)
		case *ast.FragmentSpread:
			n := s.(*ast.FragmentSpread).Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				return nil, fmt.Errorf("getSelectedFields: no fragment found with name %v", n)
			}
			sel, err := selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)
			if err != nil {
				return nil, err
			}
			selected = append(selected, sel...)
		default:
			return nil, fmt.Errorf("getSelectedFields: found unexpected selection type %v", t)
		}
	}
	return selected, nil
}

func GetColumnListFromAttributes(model interface{}, selectedFields []string) string {

	attr := reflect.ValueOf(model)
	var selectedColumns []string

	for i := 0; i < attr.NumField(); i++ {
		fieldType := attr.Field(i).Kind()
		switch fieldType {
		case reflect.Slice:
			continue
		case reflect.Struct:
			continue
		default:
			jsonVal := attr.Type().Field(i).Tag.Get("json")
			gormVal := attr.Type().Field(i).Tag.Get("gorm")
			column := strings.Split(gormVal, ":")[1] // make sure gormVal is >>> gorm:"column:column_name"

			if helper.IsAvailable(selectedFields, jsonVal) {
				selectedColumns = append(selectedColumns, column)
			}
		}
	}

	result := strings.Join(selectedColumns, ",")
	return strings.TrimRight(result, ",")
}
