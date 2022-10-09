function change_state(this: Entity<EntityTypes>, state: string) {
    if ((state in this.states) === false) {
        throw new Error(`State "${state}" does not exists.`)

    }

    this.previous_states.push([this.current_state, this.elapsed_states])

    this.elapsed_states = 0
    this.current_state = state
}

// ? maybe I should use a more traditional method.
// ? if this.isAttacking() { ... }
// ? if this.isMoving() { ... }
const entity_states = {
    spawn: function (this: Entity<EntityTypes>) {
        this.change_state('idle')
    },

    idle: function (this: Entity<EntityTypes>) {
        for (const str of ['jump', 'crouch', 'move'])
            this.entity.actions.includes(str) &&
            this.change_state(str)

        // Draw sprite...

        // () => requestAnimationFrame(game), 1000 / 60 = 16.66...7

        // Use "this.elapsed_states" to draw the correct sprite position.
        // "this.elapsed_states" is the amount of time the current state has been called.

        // amount: 5,
        // hold = 10
        // this.elapsed_states

        // const sprite = this.sprites[this.current_state]

        // if (this.elapsed_states >= sprite.hold * sprite.amount) {
        //    this.elapsed_states = 0
        // }

    },

    move: function (this: Entity<EntityTypes>) {
        for (const str of ['jump', 'crouch'])
            this.entity.actions.includes(str) &&
            this.change_state(str)

        if (!this.entity.actions.includes('move'))
            this.call_state('idle')
        
        // Handle movement...

        // Draw sprite...
    },

    jump: function (this: Entity<EntityTypes>) {
        // Handle jump...

        // Draw sprite...
    },

    crouch: function (this: Entity<EntityTypes>) {
        // Handle crouch...

        // Draw sprite...
    },

    attack: function (this: Entity<EntityTypes>) {
        // Handle attack...

        // Draw sprite...

        this.set_previous_state()
    },

    death: function (this: Entity<EntityTypes>) {
        // Handle death...

        // Draw sprite...
    }
}

//

function Entity_ <T extends EntityTypes> (this: Entity<T>, entity: T) {
    const image_src = 'assets/soldier/' + entity.color

    this.sprites = {
        'idle': {
            image_src: image_src + 'idle.png',
            amount: 5,
            hold: 0 // This is to keep the frame for a certain time.
        },
        'move': {
            image_src: image_src + 'move.png',
            amount: 6,
            hold: 0
        },
        'jump': {
            image_src: image_src + 'jump.png',
            amount: 2,
            hold: 0
        },
        'crouch': {
            image_src: image_src + 'crouch.png',
            amount: 3,
            hold: 0
        },
        'death': {
            image_src: image_src + 'death.png',
            amount: 8,
            hold: 0
        }
    }

}

const Entity = Entity_ as unknown as {
    new (entity: any): {}
}

export { Entity }
