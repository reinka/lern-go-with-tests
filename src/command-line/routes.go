package poker

func (p *PlayerServer) rountes() {
	p.router.Handle("/league", p.handleLeague())
	p.router.Handle("/players/", p.handlePlayers())

}
