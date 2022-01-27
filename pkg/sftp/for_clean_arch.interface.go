package sftp

type IPoint interface {
	Upload(fromPath, toPath string) error
}
