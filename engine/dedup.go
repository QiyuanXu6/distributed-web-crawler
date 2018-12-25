package engine

type DedupService struct {
	seen map[string]bool
}

func NewDedupService()  *DedupService {
	return &DedupService{
		seen: make(map[string]bool),
	}
}

func (d *DedupService) isDup(url string) bool {
	if d.seen[url] {
		return true
	}
	d.seen[url] = true
	return false
}