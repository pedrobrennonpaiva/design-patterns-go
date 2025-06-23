package bridges

type Device interface {
	getName() string
	setPower(power bool)
	getPower() bool
	setVolume(volume int)
	getVolume() int
}

type TV struct {
	power  bool
	volume int
	name   string
}

func NewTV() *TV {
	return &TV{
		power:  false,
		volume: 10,
		name:   "TV",
	}
}

func (t *TV) getName() string {
	return t.name
}

func (t *TV) setPower(power bool) {
	t.power = power
}

func (t *TV) getPower() bool {
	return t.power
}

func (t *TV) setVolume(volume int) {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	t.volume = volume
}

func (t *TV) getVolume() int {
	return t.volume
}

type Radio struct {
	power  bool
	volume int
	name   string
}

func NewRadio() *Radio {
	return &Radio{
		power:  false,
		volume: 1,
		name:   "Radio",
	}
}

func (t *Radio) getName() string {
	return t.name
}

func (t *Radio) setPower(power bool) {
	t.power = power
}

func (t *Radio) getPower() bool {
	return t.power
}

func (t *Radio) setVolume(volume int) {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	t.volume = volume
}

func (t *Radio) getVolume() int {
	return t.volume
}
