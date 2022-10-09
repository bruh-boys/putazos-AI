type EntityTypes = Soldier | Projectile

interface EntityMethods {
    change_state: (this: Entity<EntityTypes>, state: string) => void,
    set_previous_state: (this: Entity<EntityTypes>) => void,
}

interface Position {
    x: number,
    y: number,
    z: number
}

interface State {
    current: string,
    name: string
}

interface Entity <T extends EntityTypes> extends EntityMethods {
    position: Position,

    states: {
        [key: string]: (e: Entity<T>, i: number) => void
    },

    prev_states: State[],
    state: State
}
