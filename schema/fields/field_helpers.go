package fields

import "github.com/graphql-go/graphql"

// GenArg 简化gql参数定义
func GenArg(gqlType graphql.Input, des string, opts ...interface{}) *graphql.ArgumentConfig {
  defaultValue := interface{}(nil)
  if len(opts) > 0 && !opts[0].(bool) {
    gqlType = graphql.NewNonNull(gqlType)
  }

  if len(opts) > 1 {
    defaultValue = opts[1]
  }

  return &graphql.ArgumentConfig{
    Type:         gqlType,
    Description:  des,
    DefaultValue: defaultValue,
  }
}
