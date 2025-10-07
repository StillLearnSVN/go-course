package handlers

import (
	"errors"
	"net/http"
	"reflect"
	"restapi/pkg/utils"
	"strings"
)

func CheckBlankFields(value interface{}) error {
	val := reflect.ValueOf(value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			return utils.ErrorHandler(errors.New("all fields are required"), "All fields are required")
		}
	}
	return nil
}

func GetFieldNames(model interface{}) []string {
	val := reflect.TypeOf(model)
	fields := []string{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldToAdd := strings.TrimSuffix(field.Tag.Get("json"), ",omitempty")
		fields = append(fields, fieldToAdd)
	}
	return fields
}

func authorizeRequest(w http.ResponseWriter, r *http.Request, allowedRoles ...string) (string, bool) {
	roleVal := r.Context().Value(utils.ContextKey("role"))
	role, ok := roleVal.(string)
	if !ok || role == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return "", false
	}

	if len(allowedRoles) == 0 {
		return role, true
	}

	if _, err := utils.AuthorizeUser(role, allowedRoles...); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return "", false
	}

	return role, true
}

func userIDFromContext(r *http.Request) (int, bool) {
	val := r.Context().Value(utils.ContextKey("userId"))
	switch v := val.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}
