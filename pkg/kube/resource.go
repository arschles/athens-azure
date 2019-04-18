package kube

type Resource interface {
	Installer
	Updater
	Getter
}
