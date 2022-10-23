class Entity implements IEntity {
    protected readonly ctx: CanvasRenderingContext2D

    // @ts-ignore - Value inserted by the child class.
    protected readonly height: number

    // @ts-ignore - Value inserted by the child class.
    protected readonly width: number

    // @ts-ignore - Value inserted by the child class.
    protected readonly sprites: { [key: string]: {
        image_src: string,
        amount: number,
        hold: number
    } }

    // @ts-ignore - Value inserted by the child class.
    protected readonly states: {
        [key: string]: (this: Entity) => void
    }

    protected previous_states: {
        elapsed: number,
        state: string
    }[] = []

    protected elapsed_states: number = 0
    protected current_state: string = 'spawn'

    public direction: boolean = true

    public position: Position = {
        x: 0, y: 0, z: 0
    }

    public velocity: Position = {
        x: 0, y: 0, z: 0
    }

    // Like a muzzle flash of a gun.
    protected hidden_actions: string[] = []

    public cosmetics: string[] = []
    public actions: string[] = []

    constructor(ctx: CanvasRenderingContext2D, entity: IEntity) {
        for (const [key, value] of Object.entries(entity))
            (this as any)[key] = value

        this.ctx = ctx
    }

    public reset_state_if_ended(): boolean {
        const { amount, hold } = this.sprites[this.current_state]

        if (this.elapsed_states > amount * hold)
            return false
    
        this.change_state(this.current_state)

        return true
    }

    public includes_some_action(actions: string[]) {
        for (const action of this.actions)
            for (const act of actions) if (action === act)
                return true

        return false
    }

    public change_state_if_action(actions: string[]) {
        if (!this.includes_some_action(actions))
            return false
        
        this.change_state(actions[0])
        return true
    }

    public change_state_to_previous() {
        const { elapsed, state } = this.previous_states.pop() || 
            { elapsed: 0, state: 'spawn' }

        this.elapsed_states = elapsed
        this.current_state = state
    }

    public change_state(state: string) {
        this.previous_states.push({
            elapsed: this.elapsed_states,
            state: this.current_state
        })

        this.elapsed_states = 0
        this.current_state = state
    }

    protected update_() {
        while (true) {
            const previous_state = this.current_state,
                previos_elapsed = this.elapsed_states

            if (this.current_state in this.states)
                this.states[this.current_state].call(this)
            
            else {} // Remove this entity.

            if (
                this.elapsed_states !== previos_elapsed ||
                this.current_state !== previous_state
            ) continue

            this.elapsed_states++
            break
        }

    }

    public update() {}
}

export { Entity }
