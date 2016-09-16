package object

// Object basic object stucture
type Object struct {
    name string
    description string
}

type Creature struct {
    Object
    age int
    species string
    gender string //TODO change to male/female/unknown enum
}

type Player struct {
    Creature
    title string
    alignment string //TODO enum for alignment DnD
    equipment map[string]string
}

type Monster struct {
    Creature
    alignment string //change to DnD style alignment enum
}
