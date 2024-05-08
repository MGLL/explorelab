package queue

type Queue interface {
	PublishPlainText([]byte)
}
