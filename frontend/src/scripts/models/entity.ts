class Entity implements IEntity {
    protected readonly ctx: CanvasRenderingContext2D

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

    public position: IPosition = {
        x: 0, y: 0, z: 0
    }

    public velocity: IPosition = {
        x: 0, y: 0, z: 0
    }

    public actions: string[] = []

    constructor(ctx: CanvasRenderingContext2D, entity: IEntity) {
        for (const [key, value] of Object.entries(entity))
            (this as any)[key] = value

        this.ctx = ctx
    }

    public change_state_to_previous() {
        
    }

    public change_state(state: string) {

    }

    protected update_() {
        while (true) {
            const previous_state = this.current_state

            if (this.current_state in this.states)
                this.states[this.current_state].call(this)

            if (this.current_state !== previous_state)
                continue

            this.elapsed_states++
            break
        }

    }

    public update() {}
}

export { Entity }
