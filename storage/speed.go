package storage

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const localStorageSpeed = "speed"

// Speed represents the games speed.
type Speed string

// List of possible difficulties.
const (
	Slow   Speed = "slow"
	Normal Speed = "normal"
	Fast   Speed = "fast"
)

// LoadSpeed loads the speed setting from local storage.
// If the value is not supported, it defaults to [Normal]
func LoadSpeed(ctx app.Context) Speed {
	var val string
	ctx.LocalStorage().Get(localStorageSpeed, &val)
	d := Normal
	switch val {
	case string(Slow):
		d = Slow
	case string(Normal):
		d = Normal
	case string(Fast):
		d = Fast
	}
	return d
}

// SetSpeed stores the provided value. If the value is not supported,
// it defaults to [Normal]
func SetSpeed(ctx app.Context, val string) Speed {
	d := Normal
	switch val {
	case string(Slow):
		d = Slow
	case string(Normal):
		d = Normal
	case string(Fast):
		d = Fast
	}
	ctx.LocalStorage().Set(localStorageSpeed, d)
	return d
}
