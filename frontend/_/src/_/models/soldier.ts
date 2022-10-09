const soldier_states: Entity<Soldier>['states'] = {
    'spawn': (e: Entity<Soldier>, i: number) => {
        
    },
    'idle': (e: Entity<Soldier>, i: number) => {},
    'move': (e: Entity<Soldier>, i: number) => {},
    'attack': (e: Entity<Soldier>, i: number) => {},
    'die': (e: Entity<Soldier>, i: number) => {},
}

function _Soldier(this: Entity<Soldier>, entity: SimplifiedEntity<Soldier>) {
    this.position = entity.position
    this.entity = entity.entity

    this.states = soldier_states
}

const Soldier = _Soldier as unknown as {
    new (entity: SimplifiedEntity<Soldier>): Entity<Soldier>
}
