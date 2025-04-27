package agenda

import "time"

type Agenda struct {
	Id           uint64        `gorm:"primarykey"`
	Nome         string        `gorm:"size:50;not null"`
	Descricao    *string       `gorm:"size:100"`
	Apontamentos []Apontamento `gorm:"foreignKey:IdAgenda"`
}

type Apontamento struct {
	Id        uint64    `gorm:"primarykey"`
	Inicio    time.Time `gorm:"index;not null"`
	Fim       time.Time `gorm:"not null"`
	IdAgenda  uint64
	IdCliente uint64
	Cliente   Cliente `gorm:"foreignKey:IdCliente"`
	Status    string  `gorm:"size:1;not null;default:'A'"`
}

type Cliente struct {
	Id           uint64        `gorm:"primarykey"`
	Nome         string        `gorm:"size:50;not null"`
	Apontamentos []Apontamento `gorm:"foreignKey:IdCliente"`
}
