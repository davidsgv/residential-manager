package entities

import "slices"

type Permission []string

// check if the permissions in the argument exists in the type
func (permission *Permission) ArePermissionsPresent(p2 Permission) bool {
	for _, e := range p2 {
		if !slices.Contains(*permission, e) {
			return false
		}
	}
	return true
}
