import { Entity } from './@module.js'

const soldier_states: Entity['states'] = {
    spawn: function () {
        this.change_state('idle')

    },

    idle: function () {
        if (this.actions.includes('move'))
            return this.change_state('walk')

    },

    move: function () {
        if (!this.actions.includes('move'))
            return this.change_state('idle')

        this.direction === true
            ? this.position['x'] += this.velocity['x']
            : this.position['x'] -= this.velocity['x']

        if (this.actions.includes('jump'))
            return this.change_state('jump')
    
    },

    jump: function () {
        if (!this.actions.includes('jump'))
            return this.change_state('fall')

        // Handle jump
    },

    fall: function () {
        // Check if the entity is on the ground.
        
    },

    death: function () {
        
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
