package enum

type State int

type EnumVariable struct {
	Key   State
	Value string
}

type Enum map[State]string

func Create(evs []EnumVariable) Enum {

	enum := make(Enum)

	for ev := range evs {
		enum[evs[ev].Key] = evs[ev].Value
	}

	return enum
}
