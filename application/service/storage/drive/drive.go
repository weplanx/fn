package drive

type API interface {
	Put(filename string, body []byte) (err error)
}
