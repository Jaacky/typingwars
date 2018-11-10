package typingwars

import (
	"log"

	"github.com/gofrs/uuid"
)

// Game strcut
type Game struct {
	Clients         map[uuid.UUID]*Client
	Teams           []*Team
	Space           *Space
	InGame          bool
	EventDispatcher *EventDispatcher
	physicsTicker   *PhysicsTicker
	unitSpawner     *UnitSpawner
}

// NewGame struct
func NewGame(room *Room) *Game {
	// bases := []*baseBuilding{}
	clients := room.clients
	teams := makeTeams(clients, 2)
	space := NewSpace(clients)

	eventDispatcher := NewEventDispatcher()
	physicsTicker := NewPhysicsTicker(eventDispatcher)
	unitSpawner := NewUnitSpawner(eventDispatcher, space, teams)

	updater := NewUpdater(space, eventDispatcher)
	eventDispatcher.RegisterTimeTickListener(updater)
	eventDispatcher.RegisterUnitSpawnedListener(updater)
	eventDispatcher.RegisterUserActionListener(updater)
	eventDispatcher.RegisterPhysicsReadyListener(room)

	return &Game{
		Space:           space,
		Teams:           teams,
		Clients:         clients,
		EventDispatcher: eventDispatcher,
		physicsTicker:   physicsTicker,
		unitSpawner:     unitSpawner,
	}
}

func makeTeams(clients map[uuid.UUID]*Client, numTeams int) []*Team {
	teams := []*Team{}

	for i := 0; i < numTeams; i++ {
		teams = append(teams, NewTeam())
		log.Printf("Team %d made!", i)
	}

	j := 0
	for _, client := range clients {
		teamNum := j % numTeams
		log.Printf("Client to team %d!", teamNum)
		teams[teamNum].AddPlayer(client)
		j++
	}

	return teams
}

func (g *Game) start() {
	go g.EventDispatcher.RunEventLoop()
	go g.physicsTicker.Run()
	go g.unitSpawner.Run()
}
