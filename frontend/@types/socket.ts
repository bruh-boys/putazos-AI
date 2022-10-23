type SocketValue <T extends 'join' | 'data'> =
    T extends 'join' ? World : {
        projectiles: EntityModel<ProjectileModel>[],
        soldiers: EntityModel<SoldierModel>[]
    }

interface SocketResponse {
    type: 'join' | 'data',
    data: SocketValue<SocketResponse['type']>
}

interface SocketRequest {
    type: 'action',
    data: {
        state: boolean,
        key: string
    }
}
