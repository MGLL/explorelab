package manager

type Queue interface {
	PublishPlainText([]byte)
	GetName() string
}
