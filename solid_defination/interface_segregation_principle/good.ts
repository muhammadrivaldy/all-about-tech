interface PondBigFish {
    catfish(): void
    swordfish(): void
}

interface PondSmallFish {
    salmon(): void
}

class BigFish implements PondBigFish {

    public catfish(): void {
        // do something
    }

    public swordfish(): void {
        // do something
    }

}

class SmallFish implements PondSmallFish {

    public salmon(): void {
        // do something
    }

}