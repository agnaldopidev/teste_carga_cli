package domain

type HTTPClient interface {
	Get(url string) (int, error)
}
