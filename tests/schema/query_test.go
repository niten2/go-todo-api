package testing

import (
  // "fmt"
  "testing"
	"reflect"
  "github.com/joho/godotenv"
  "github.com/graphql-go/graphql"
  . "github.com/smartystreets/goconvey/convey"

  . "go-todo-api/models"
  "go-todo-api/db"
  . "go-todo-api/graphql"
	// "gopkg.in/mgo.v2/bson"
)

func init() {
  _ = godotenv.Load("../../.env.test")
	db.Connect()
}

func TestQuery(t *testing.T) {
  DbProduct := db.Db.C("products")
  DbProduct.RemoveAll(nil)

  Convey("AllProduct", t, func() {
    CreateProduct(Product{Id: "1", Name: "test1"})
    CreateProduct(Product{Id: "2", Name: "test2"})

    query := "query { products { id name } }"
    res := ExecuteQuery(query)

    expected := &graphql.Result{
      Data: map[string]interface{}{
        "products": []interface{}{
          map[string]interface{}{
            "id":   "1",
            "name": "test1",
          },
          map[string]interface{}{
            "id":   "2",
            "name": "test2",
          },
        },
      },
    }

    eq := reflect.DeepEqual(res, expected)
    So(eq, ShouldBeTrue)
  })
}

// package testing

// import (
// 	// "context"
//   // "fmt"
// 	"reflect"
// 	"testing"
// 	"github.com/graphql-go/graphql"
// 	"github.com/graphql-go/graphql/testutil"

//   myGraphql "go-todo-api/graphql"
// )

// // func init () {
// //   fmt.Println(graphql.Schema)
// // }

// type T struct {
// 	Query    string
// 	Schema   graphql.Schema
// 	Expected interface{}
// }

// var Tests = []T{}

// func init() {
// 	Tests = []T{
// 		{
// 			Query: `
// 				query products {
// 					products {
// 						name
// 					}
// 				}
// 			`,
// 			Schema: myGraphql.Schema,
// 			Expected: &graphql.Result{
// 				Data: map[string]interface{}{
// 					"hero": map[string]interface{}{
// 						"name": "R2-D2",
// 					},
// 				},
// 			},
// 		},

// 		// {
// 		// 	Query: `
// 		// 		query HeroNameAndFriendsQuery {
// 		// 			hero {
// 		// 				id
// 		// 				name
// 		// 				friends {
// 		// 					name
// 		// 				}
// 		// 			}
// 		// 		}
// 		// 	`,
// 		// 	Schema: testutil.StarWarsSchema,
// 		// 	Expected: &graphql.Result{
// 		// 		Data: map[string]interface{}{
// 		// 			"hero": map[string]interface{}{
// 		// 				"id":   "2001",
// 		// 				"name": "R2-D2",
// 		// 				"friends": []interface{}{
// 		// 					map[string]interface{}{
// 		// 						"name": "Luke Skywalker",
// 		// 					},
// 		// 					map[string]interface{}{
// 		// 						"name": "Han Solo",
// 		// 					},
// 		// 					map[string]interface{}{
// 		// 						"name": "Leia Organa",
// 		// 					},
// 		// 				},
// 		// 			},
// 		// 		},
// 		// 	},
// 		// },

// 	}
// }

// func TestQuery(t *testing.T) {
//   // fmt.Println(Schema)

// 	for _, test := range Tests {
// 		params := graphql.Params{
// 			Schema:        test.Schema,
// 			RequestString: test.Query,
// 		}
// 		testGraphql(test, params, t)
// 	}
// }

// func testGraphql(test T, p graphql.Params, t *testing.T) {
// 	result := graphql.Do(p)

// 	if len(result.Errors) > 0 {
// 		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
// 	}

// 	if !reflect.DeepEqual(result, test.Expected) {
// 		t.Fatalf("wrong result, query: %v, graphql result diff: %v", test.Query, testutil.Diff(test.Expected, result))
// 	}
// }



// func TestBasicGraphQLExample(t *testing.T) {
// 	// taken from `graphql-js` README

// 	helloFieldResolved := func(p graphql.ResolveParams) (interface{}, error) {
// 		return "world", nil
// 	}

// 	schema, err := graphql.NewSchema(graphql.SchemaConfig{
// 		Query: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "RootQueryType",
// 			Fields: graphql.Fields{
// 				"hello": &graphql.Field{
// 					Description: "Returns `world`",
// 					Type:        graphql.String,
// 					Resolve:     helloFieldResolved,
// 				},
// 			},
// 		}),
// 	})
// 	if err != nil {
// 		t.Fatalf("wrong result, unexpected errors: %v", err.Error())
// 	}
// 	query := "{ hello }"
// 	var expected interface{}
// 	expected = map[string]interface{}{
// 		"hello": "world",
// 	}

// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})
// 	if len(result.Errors) > 0 {
// 		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
// 	}
// 	if !reflect.DeepEqual(result.Data, expected) {
// 		t.Fatalf("wrong result, query: %v, graphql result diff: %v", query, testutil.Diff(expected, result))
// 	}

// }

// func TestThreadsContextFromParamsThrough(t *testing.T) {
// 	extractFieldFromContextFn := func(p graphql.ResolveParams) (interface{}, error) {
// 		return p.Context.Value(p.Args["key"]), nil
// 	}

// 	schema, err := graphql.NewSchema(graphql.SchemaConfig{
// 		Query: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "Query",
// 			Fields: graphql.Fields{
// 				"value": &graphql.Field{
// 					Type: graphql.String,
// 					Args: graphql.FieldConfigArgument{
// 						"key": &graphql.ArgumentConfig{Type: graphql.String},
// 					},
// 					Resolve: extractFieldFromContextFn,
// 				},
// 			},
// 		}),
// 	})
// 	if err != nil {
// 		t.Fatalf("wrong result, unexpected errors: %v", err.Error())
// 	}
// 	query := `{ value(key:"a") }`

// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 		Context:       context.WithValue(context.TODO(), "a", "xyz"),
// 	})
// 	if len(result.Errors) > 0 {
// 		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
// 	}
// 	expected := map[string]interface{}{"value": "xyz"}
// 	if !reflect.DeepEqual(result.Data, expected) {
// 		t.Fatalf("wrong result, query: %v, graphql result diff: %v", query, testutil.Diff(expected, result))
// 	}

// }

// func TestNewErrorChecksNilNodes(t *testing.T) {
// 	schema, err := graphql.NewSchema(graphql.SchemaConfig{
// 		Query: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "Query",
// 			Fields: graphql.Fields{
// 				"graphql_is": &graphql.Field{
// 					Type: graphql.String,
// 					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 						return "", nil
// 					},
// 				},
// 			},
// 		}),
// 	})
// 	if err != nil {
// 		t.Fatalf("unexpected errors: %v", err.Error())
// 	}
// 	query := `{graphql_is:great(sort:ByPopularity)}{stars}`
// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})
// 	if len(result.Errors) == 0 {
// 		t.Fatalf("expected errors, got: %v", result)
// 	}
// }

// func TestEmptyStringIsNotNull(t *testing.T) {
// 	checkForEmptyString := func(p graphql.ResolveParams) (interface{}, error) {
// 		arg := p.Args["arg"]
// 		if arg == nil || arg.(string) != "" {
// 			t.Errorf("Expected empty string for input arg, got %#v", arg)
// 		}
// 		return "yay", nil
// 	}
// 	returnEmptyString := func(p graphql.ResolveParams) (interface{}, error) {
// 		return "", nil
// 	}

// 	schema, err := graphql.NewSchema(graphql.SchemaConfig{
// 		Query: graphql.NewObject(graphql.ObjectConfig{
// 			Name: "Query",
// 			Fields: graphql.Fields{
// 				"checkEmptyArg": &graphql.Field{
// 					Type: graphql.String,
// 					Args: graphql.FieldConfigArgument{
// 						"arg": &graphql.ArgumentConfig{Type: graphql.String},
// 					},
// 					Resolve: checkForEmptyString,
// 				},
// 				"checkEmptyResult": &graphql.Field{
// 					Type:    graphql.String,
// 					Resolve: returnEmptyString,
// 				},
// 			},
// 		}),
// 	})
// 	if err != nil {
// 		t.Fatalf("wrong result, unexpected errors: %v", err.Error())
// 	}
// 	query := `{ checkEmptyArg(arg:"") checkEmptyResult }`

// 	result := graphql.Do(graphql.Params{
// 		Schema:        schema,
// 		RequestString: query,
// 	})
// 	if len(result.Errors) > 0 {
// 		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
// 	}
// 	expected := map[string]interface{}{"checkEmptyArg": "yay", "checkEmptyResult": ""}
// 	if !reflect.DeepEqual(result.Data, expected) {
// 		t.Errorf("wrong result, query: %v, graphql result diff: %v", query, testutil.Diff(expected, result))
// 	}
// }
