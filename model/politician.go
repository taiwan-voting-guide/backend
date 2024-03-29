package model

import (
	"time"
)

type Politician struct {
	Id        int
	Name      string
	Birthdate time.Time
	AvatarUrl string
	Sex       Sex
	Meta      []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Politician) Repr() *PoliticianRepr {
	return &PoliticianRepr{
		Name:      p.Name,
		Birthdate: p.Birthdate,
		AvatarUrl: p.AvatarUrl,
		Sex:       p.Sex,

		CreatedAt: p.CreatedAt.Unix(),
		UpdatedAt: p.UpdatedAt.Unix(),
	}
}

type PoliticianRepr struct {
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	AvatarUrl string    `json:"avatarUrl"`
	Sex       Sex       `json:"sex"`

	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

type Sex string

const (
	SexMale   = "male"
	SexFemale = "female"
)
