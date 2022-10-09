package types

type Soldier struct {
	Actions []string
	Color   string
}

func (s Soldier) HasAction(ss string) bool {
	for _, a := range s.Actions {
		if a == ss {

			return true
		}

	}

	return false
}

type World struct{}

func (w World) GenerateEntity(s string) bool {

	return false
}

type Entity[T Soldier] struct {
	ID string

	Sprites   map[string]string
	States    map[string]func()
	PrevState string
	State     string

	World  *World
	Entity *T
}

func (e *Entity[any]) ChangeState(state string) bool {
	if _, ok := e.States[state]; ok == true {
		e.PrevState = e.State
		e.State = state

		return true
	}

	return false
}

func (e *Entity[any]) SetPreviosState() {
	e.State = e.PrevState
}

func Contains[T string](s T, ss []T) bool {
	for _, a := range ss {
		if s == a {
			return true

		}

	}

	return false
}

var SoldierStates = map[string]func(e Entity[Soldier]){
	"spawn": func(e Entity[Soldier]) {
		e.ChangeState("idle")
	},
	"idle": func(e Entity[Soldier]) {
		for _, action := range e.Entity.Actions {
			if Contains(action, []string{"move", "attack"}) {
				e.ChangeState(action)

				break
			}

		}

	},
	"attack": func(e Entity[Soldier]) {
		e.World.GenerateEntity("bullet")
		e.SetPreviosState()
	},
}