package bin

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func (b *Bin) NewBin() *Bin {
	return &Bin{
		Id:        b.Id,
		Private:   b.Private,
		CreatedAt: b.CreatedAt,
		Name:      b.Name,
	}
}
