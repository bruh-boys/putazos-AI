interface IEntity {
    direction: boolean, // false = left, true = right
    position: IPosition,
    velocity: IPosition,
    actions: string[]
}

interface ISoldier extends IEntity {
    color: string,
    id: string,
}
