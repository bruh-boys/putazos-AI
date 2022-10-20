interface Position {
    x: number, y: number, z: number
}

interface Def {
    radius: Position,
    positions: Position[],
    image_src: string,
    visible: boolean,
    id: string,
}

interface World {
    radius: Position
    defs: Def[]
}
