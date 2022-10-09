function set_previous_state(this: Entity<EntityTypes>, state: string) {
    this.state = this.prev_states.pop() || this.state

}

function change_state(this: Entity<EntityTypes>, state: string) {
    if (!Object.keys(this.states).includes(state)) {
        throw new Error(`State "${state}" does not exists.`)

    }

    this.prev_states.push(this.state)
    this.state = state
}
