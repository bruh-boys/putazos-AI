import { Entity } from './@module.js'

const soldier_states: {
    [key: string]: (this: Soldier) => void
} = {
    spawn: function (this: Soldier) {
        return this.change_state('idle')
    },

    idle: function (this: Soldier) {
        // In case the state is in the last frame, reset it.
        if (this.reset_state_if_ended()) return

        if (this.change_state_if_action(
            [ 'attack', 'move', 'fall', 'jump' ]
        )) return

        // Draw sprite...
    },

    move: function (this: Soldier) {
        // In case the state is in the last frame, reset it.
        if (this.reset_state_if_ended()) return

        if (!(
            this.includes_some_action(['fall', 'jump']) &&
            this.actions.includes('move')
        ))
            return this.change_state('idle')

        // ? I think isn't necessary in the frontend.
        // ? The client will send the position to the server
        // ? Or the server will send the position to the client.
        // ? Maybe create a new sub function to handle this.
        this.direction === true
            ? this.position['x'] += this.velocity['x']
            : this.position['x'] -= this.velocity['x']

        if (this.change_state_if_action(
            [ 'attack', 'idle', 'fall', 'jump' ]
        )) return

        // Draw sprite...
    },

    jump: function (this: Soldier) {
        if (this.change_state_if_action(
            [ 'attack', 'move' ]
        )) return

        // Handle jump
    },

    fall: function (this: Soldier) {
        if (this.change_state_if_action(
            [ 'attack', 'move' ]
        )) return

        // Check if the entity is on the ground.

    },

    attack: function (this: Soldier) {
        // Handle attack

        // Add cosmetics, like a muzzle flash of a gun.
        this.cosmetics.push('muzzle_flash')
    },

    death: function (this: Soldier) {
        // Handle death

        // Draw sprite...
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

        // I need check width and height of the sprite.
        // this.height = 0
        // this.width = 0
    }

    public update() {
        this.update_()

    }

}

export { Soldier }
