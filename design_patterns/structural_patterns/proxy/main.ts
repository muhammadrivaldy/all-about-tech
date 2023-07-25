interface Humans {
    setDo(s: string): void
    operation(): void
}

class Security implements Humans {
    private do: string = ""
    private human: Humans

    public constructor(human: Humans) {
        this.human = human
    }

    public setDo(s: string): void {
        this.do = s
    }

    public operation(): void {
        this.human.operation()
        console.log(this.do)
    }
}

class Interviewer implements Humans {
    private do: string = ""
    
    public setDo(s: string): void {
        this.do = s
    }

    public operation(): void {
        console.log(this.do)
    }
}

var interviewer = new Interviewer()
interviewer.setDo("Interviewer wants do interview")

var securityLobby = new Security(interviewer)
securityLobby.setDo("Security lobby accompanies interviewer to the target floor")

var securityFloor = new Security(securityLobby)
securityFloor.setDo("Security lobby accompanies security lobby and interviewer to the target room")

securityFloor.operation()