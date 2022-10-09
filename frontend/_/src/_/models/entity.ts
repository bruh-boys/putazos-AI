function set_previous_state(this: Entity<EntityTypes>) {
    this.state = this.prev_states.pop() || this.state

}

function change_state(this: Entity<EntityTypes>, state: string) {
    if (!Object.keys(this.states).includes(state)) {
        throw new Error(`State "${state}" does not exists.`)

    }

    this.prev_states.push(this.state)
    this.state = [state, 0]
}

function change_position(
    this: Entity<EntityTypes>, pos: { x?: number, y?: number, z?: number }
) {
    for (const axis of Object.keys(pos)) if (pos[axis as 'x'] !== undefined)
        this.position[axis as 'x'] = pos[axis as 'x']!

}

