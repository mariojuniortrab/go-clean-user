package protocol_entity

type Emptyable interface {
	IsEmpty() bool
	New() Emptyable
}
