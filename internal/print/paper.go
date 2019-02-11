package print

import "github.com/hasandi/print-packer/internal/shape"

type Paper struct {
	length, width float64
}

func NewPaper(length, width float64) *Paper {
	return &Paper{length: length, width: width}
}

func (p *Paper) Rotate90Deg() {
	p.length, p.width = p.width, p.length
}

func (p Paper) Length() float64 {
	return p.length
}

func (p Paper) Width() float64 {
	return p.width
}

func (p Paper) IsSquare() bool {
	return p.length == p.width
}

func (p Paper) IsPortrait() bool {
	if p.IsSquare() {
		return true
	}
	return p.length < p.width
}

func (p Paper) IsLandscape() bool {
	if p.IsSquare() {
		return true
	}
	return p.width < p.length
}

func (p *Paper) Portrait() *Paper {
	if !p.IsPortrait() {
		p.Rotate90Deg()
	}
	return p
}

func (p *Paper) Landscape() *Paper {
	if !p.IsLandscape() {
		p.Rotate90Deg()
	}
	return p
}

func (p Paper) Orientations() []*Paper {
	portrait := NewPaper(p.Portrait().Length(), p.Portrait().Width())
	landscape := NewPaper(p.Landscape().Length(), p.Landscape().Width())

	return []*Paper{portrait, landscape}
}

func (p Paper) PackTo(obj shape.Rectangle) []*Paper {
	var res []*Paper

	for _, o := range p.Orientations() {
		for i := o.Length(); i <= obj.Length(); i += o.Length() {
			for j := o.Width(); j <= obj.Width(); j += o.Width() {
				res = append(res, NewPaper(i, j))
			}
		}
	}

	return res
}
