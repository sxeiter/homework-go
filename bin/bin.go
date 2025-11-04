package bin

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func (b *Bin) NewBin(name string, private bool) *Bin {
	return &Bin{
		Id:        b.Id,
		Private:   b.Private,
		CreatedAt: b.CreatedAt,
		Name:      b.Name,
	}
}
