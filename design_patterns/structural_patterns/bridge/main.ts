interface FeaturesInterface {
    TractionControl(): FeaturesInterface
    RidingMode(): FeaturesInterface
    SmartKey(): FeaturesInterface
    ABS(): FeaturesInterface
}

class Features implements FeaturesInterface {

    public TractionControl(): FeaturesInterface {
        console.log("Traction control")
        return this
    }

    public RidingMode(): FeaturesInterface {
        console.log("Riding mode")
        return this
    }

    public SmartKey(): FeaturesInterface {
        console.log("Smart key")
        return this
    }

    public ABS(): FeaturesInterface {
        console.log("ABS")
        return this
    }

}

class MotorcycleFactory {

    private features: FeaturesInterface

    constructor() {
        this.features = new Features()
    }

    public featuresOfNMAX() {
        console.log("This is list features of NMAX:")
        this.features.ABS().SmartKey().TractionControl()
    }

    public featuresOfGenio() {
        console.log("This is list features of Genio:")
        this.features.SmartKey()
    }

}

var motorcycle = new MotorcycleFactory()
motorcycle.featuresOfGenio()
motorcycle.featuresOfNMAX()