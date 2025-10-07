package utils

import "slices"

import "errors"

type ContextKey string

func  AuthorizeUser(userRole string, allowedRoles ...string) (bool, error) {
	if slices.Contains(allowedRoles, userRole) {
			return true, nil
		}
	return false, errors.New("unauthorized: user does not have the required role")
}