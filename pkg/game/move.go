package game

type Move struct {
	startX           int
	startY           int
	endX             int
	endY             int
	pawnReachedFinal bool
}

func (m *Move) New(startX int, startY int, endX int, endY int) {
	m.startX = startX
	m.startY = startY
	m.endX = endX
	m.endY = endY
	m.pawnReachedFinal = false
}
