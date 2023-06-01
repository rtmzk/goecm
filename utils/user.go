package utils

import "os/user"

func CurrentUser() string {
	user, err := user.Current()
	if err != nil {
		return "root"
	}
	return user.Username
}

func UserHome() string {
	user, err := user.Current()
	if err != nil {
		return "root"
	}

	return user.HomeDir
}
