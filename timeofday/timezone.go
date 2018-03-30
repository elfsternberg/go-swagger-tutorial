package timeofday

type Timezone struct {
	Timezone string  `long:"timezone" short:"t" description:"The default time zone" env:"GOTIME_DEFAULT_TIMEZONE" default:"UTC"`
}
