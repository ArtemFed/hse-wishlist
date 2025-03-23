package xcommon

import (
	"fmt"
	"strings"
)

type FilterItem struct {
	Value    any
	Operator string
}

func QueryWhereAnd(selectQuery string, argsMap map[string]FilterItem) (string, []any) {
	return queryWhere(selectQuery, argsMap, "AND")
}

func QueryWhereOr(selectQuery string, argsMap map[string]FilterItem) (string, []any) {
	return queryWhere(selectQuery, argsMap, "OR")
}

func queryWhere(selectQuery string, argsMap map[string]FilterItem, argsConnector string) (string, []any) {
	var sb strings.Builder
	sb.WriteString(selectQuery)
	if len(argsMap) > 0 {
		sb.WriteString(" WHERE ")
	}
	connector := ""
	i := 0
	args := make([]any, len(argsMap))
	for k, v := range argsMap {
		if i == 1 {
			connector = fmt.Sprintf(" %s ", argsConnector)
		}
		sb.WriteString(fmt.Sprintf("%s %v %s $%d", connector, k, v.Operator, i+1))
		args[i] = v.Value
		i++
	}
	return sb.String(), args
}
