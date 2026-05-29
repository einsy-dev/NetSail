package server

type Domain struct{}

func (d *Domain) Fetch(domain string) (string, error) {
	return "", nil
}
