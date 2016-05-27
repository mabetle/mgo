package mauth

import (
	"fmt"
	"strings"

	"github.com/mabetle/mgo/mcore"
)

type ResRoles struct {
	res   string
	roles string
}

// RolePrefix role prefix ROLE_
const RolePrefix = "ROLE_"

// ResRoleMap {res,roles}
var ResRoleMap = [][]string{}

func RegAuthText(text string) error {
	LoadResRolesText(text)
	return nil
}

func RegRoles(res interface{}, roles string) error {
	_res, err := GetFuncRes(res)
	if err != nil {
		return err
	}
	AddResRoleMap(_res, roles)
	return nil
}

// QualifyRole add prefix to role string
func QualifyRole(role string) string {
	role = strings.ToUpper(role)
	role = strings.TrimSpace(role)

	if !strings.HasPrefix(role, RolePrefix) {
		role = RolePrefix + role
	}
	return role
}

// QualifyRoles roles string to array.
func QualifyRoles(roles string) []string {
	result := []string{}
	for _, role := range strings.Split(roles, ",") {
		role = strings.TrimSpace(role)
		if role == "" {
			continue
		}
		role = QualifyRole(role)
		result = append(result, role)
	}
	return result
}

func QualifyRolesStr(roles string) string {
	return strings.Join(QualifyRoles(roles), ",")
}

// CheckRoles check userRoles in need roles
func CheckRoles(needRoles, userRoles string) bool {
	// not found means no rights restrict
	if needRoles == "" {
		return true
	}
	// if has all means no rights restrict
	if strings.Contains(needRoles, "ALL") {
		return true
	}
	// if user roles null means not login yet.
	if userRoles == "" {
		//return false
	}

	needRolesA := QualifyRoles(needRoles)
	userRolesA := QualifyRoles(userRoles)

	// user has need roles
	for _, checkRole := range userRolesA {
		if strings.TrimSpace(checkRole) == "" {
			continue
		}
		checkRole = QualifyRole(checkRole)
		if mcore.NewString(checkRole).IsInArrayIgnoreCase(needRolesA) {
			return true
		}
	}
	// not found, no rights
	return false
}

// AddResRoleMap add res role map
func AddResRoleMap(res, roles string) {
	res = strings.ToLower(res)
	res = strings.TrimSpace(res)
	roles = QualifyRolesStr(roles)
	item := []string{res, roles}
	// overide exists res
	exist := false
	for index, vs := range ResRoleMap {
		rres := vs[0]
		if rres == res {
			exist = true
			ResRoleMap[index] = []string{res, roles}
		}
	}
	if !exist {
		ResRoleMap = append(ResRoleMap, item)
	}
}

func isMatch(res, checkRes string) bool {
	if res == checkRes {
		return true
	}

	rolePrefix := strings.TrimSuffix(res, "*")

	if strings.HasSuffix(res, "*") && strings.HasPrefix(checkRes, rolePrefix) {
		return true
	}
	return false
}

// getResNeedRoles
func getResNeedRoles(checkRes string) string {
	sb := mcore.NewStringBuffer()
	for _, rm := range ResRoleMap {
		if len(rm) < 2 {
			continue
		}
		res := rm[0]
		role := rm[1]
		if isMatch(res, checkRes) {
			sb.Append(role, ",")
		}
	}
	roles := sb.String()
	return strings.TrimSuffix(roles, ",")
}

// IsCanAccessRes
func IsCanAccessRes(checkRes, userRoles string) bool {
	checkRes = strings.ToLower(checkRes)
	checkRes = strings.TrimSpace(checkRes)
	needRoles := getResNeedRoles(checkRes)
	return CheckRoles(needRoles, userRoles)
}

// PrintIsCanAccessRes
func PrintIsCanAccessRes(checkRes, userRoles string, expect bool) {
	b := IsCanAccessRes(checkRes, userRoles)
	if b == expect {
		fmt.Printf("Passed\n")
		return
	}
	fmt.Printf("CheckAuth, Res:%s UserRoles: %s Result:%v Expect:%v\n", checkRes, userRoles, b, expect)
}

// PrintResRoleAuthMap print.
func PrintResRoleAuthMap() {
	fmt.Printf("***ResRoleAuth Config***\n")
	for _, rm := range ResRoleMap {
		if len(rm) < 2 {
			continue
		}
		fmt.Printf("\t%s : %s\n", rm[0], rm[1])
	}
}

func LoadResRolesText(text string) error {
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("no lines")
	}
	for _, line := range lines {
		ms := mcore.NewString(line).TrimSpace()
		// skip comment line or blank line or no = line
		if ms.IsHasPrefix("#") || ms.IsBlank() || !ms.IsContains("=") {
			continue
		}
		rm := ms.Split("=")

		// not a valid line
		if len(rm) != 2 {
			continue
		}

		res := strings.TrimSpace(rm[0])
		role := strings.TrimSpace(rm[1])

		AddResRoleMap(res, role)
	}
	return nil
}

// LoadAuthMapFile load.
func LoadAuthMapFile(location string) error {
	fmt.Printf("Load Res Auth Config from File: %s\n", location)
	text, err := mcore.ReadFileAll(location)
	if err != nil {
		fmt.Printf("Error Load Auth File: %s\n", err)
		return err
	}
	return LoadResRolesText(text)
}

func InitAuthMap() {
	fmt.Printf("***Init AuthMap\n")
	LoadResRolesText(defaultResRolesText)
}
