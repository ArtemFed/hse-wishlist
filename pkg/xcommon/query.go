package xcommon

import (
	"fmt"
	"strings"
)

func QueryWhereAnd(selectQuery string, argsMap map[string]any) (string, []any) {
	return queryWhere(selectQuery, argsMap, "AND")
}

func QueryWhereOr(selectQuery string, argsMap map[string]any) (string, []any) {
	return queryWhere(selectQuery, argsMap, "OR")
}

func queryWhere(selectQuery string, argsMap map[string]any, argsConnector string) (string, []any) {
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
		sb.WriteString(fmt.Sprintf("%s %v = $%d", connector, k, i+1))
		args[i] = v
		i++
	}
	return sb.String(), args
}
