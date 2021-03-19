package main

func hasPermissionOnNamespace(username string, namespace string) bool {
	return username == namespace
}