import { Entity } from './@module.js'

const soldier_states: {
    [key: string]: (this: Soldier) => void
} = {
    spawn: function (this: Soldier) {
        return this.change_state('idle')
    },

    idle: function (this: Soldier) {
        // In case the state is in the last frame, reset it.
        if (this.reset_state_if_endend()) return

        for (const action of ['attack', 'move', 'fall', 'jump']) {
            if (this.actions.includes(action))
                return this.change_state(action)
        }

        // Draw sprite...
    },

    move: function (this: Soldier) {
        // In case the state is in the last frame, reset it.
        if (this.reset_state_if_endend()) return

        if (!(
            this.actions.includes('fall') || this.actions.includes('jump') &&
            this.actions.includes('move')
        ))
            return this.change_state('idle')

        // ? I think isn't necessary in the frontend.
        // ? The client will send the position to the server
        // ? Or the server will send the position to the client.
        this.direction === true
            ? this.position['x'] += this.velocity['x']
            : this.position['x'] -= this.velocity['x']

        for (const action of ['attack', 'idle', 'fall', 'jump']) {
            if (this.actions.includes(action))
                return this.change_state(action)
        }

        // Draw sprite...
    },

    jump: function (this: Soldier) {
        for (const action of ['attack', 'move']) {
            if (this.actions.includes(action))
                return this.change_state(action)
        }

        // Handle jump
    },

    fall: function (this: Soldier) {
        for (const action of ['attack', 'move']) {
            if (this.actions.includes(action))
                return this.change_state(action)
        }

        // Check if the entity is on the ground.

    },

    attack: function (this: Soldier) {
        if (!this.actions.includes('attack'))
            return this.change_state_to_previous()

        // Handle attack
    },

    death: function (this: Soldier) {
        
    },

}

const soldier_sprites: Entity['sprites'] = {

}

class Soldier extends Entity {
    public readonly color: string = ''
    public readonly id: string = ''

    constructor(ctx: CanvasRenderingContext2D, entity: ISoldier) {
        super(ctx, entity)

        // @ts-ignore - Requerid insert.
        this.sprites = soldier_sprites

        // @ts-ignore - Requerid insert.
        this.states = soldier_states

    }

    public update() {
        this.update_()

    }

}

export { Soldier }
