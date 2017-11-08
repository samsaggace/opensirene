package conf

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Conf holds the necessary configuration for the application to work
type Conf struct {
	Database   Database   `yaml:"database"`
	Server     Server     `yaml:"server"`
	LogLevel   string     `yaml:"loglevel" env:"LOGLEVEL" default:"info"`
	Prometheus Prometheus `yaml:"prometheus"`
	Crontab    Crontab    `yaml:"crontab"`
}

// Prometheus is a simple struct holding configuration variables for prometheus
type Prometheus struct {
	Prefix string `yaml:"prefix" env:"PROMETHEUS_PREFIX" default:"opensirene"`
}

// Crontab is a simple struct holding configuration variables for the periodic
// task to be executed
type Crontab struct {
	DownloadPath string `yaml:"download_path" env:"DOWNLOAD_PATH" default:"downloads"`
	EveryXHours  uint64 `yaml:"every_x_hours" env:"EVERY_X_HOURS" default:"3"`
}

// C is the main exported configuration
var C Conf

// Parse will parse every nested fields with the env/defaults parser
// and set the values accordingly
func (c *Conf) Parse() error {
	var err error
	if err = Parse(&c.Database); err != nil {
		return errors.Wrap(err, "couldn't parse Database struct")
	}
	if err = Parse(&c.Server); err != nil {
		return errors.Wrap(err, "couldn't parse Server struct")
	}
	if err = Parse(&c.Server.Cors); err != nil {
		return errors.Wrap(err, "couldn't parse Server.Cors struct")
	}
	if err = Parse(&c.Prometheus); err != nil {
		return errors.Wrap(err, "couldn't parse Prometheus struct")
	}
	if err = Parse(&c.Crontab); err != nil {
		return errors.Wrap(err, "couldn't parse Crontab struct")
	}
	if err = Parse(c); err != nil {
		return errors.Wrap(err, "couldn't parse Conf struct")
	}
	SetLogLevel(c.LogLevel)
	if _, err = os.Stat(c.Crontab.DownloadPath); os.IsNotExist(err) {
		os.MkdirAll(c.Crontab.DownloadPath, os.ModePerm)
	}
	return Parse(c)
}

// SetLogLevel sets the logging level when possible, otherwise it fallbacks to
// the default logrus level and logs a warning
func SetLogLevel(lvl string) {
	l, err := logrus.ParseLevel(lvl)
	if err != nil {
		logrus.WithField("provided", lvl).Warn("Invalid log level, fallback to Info level")
	} else {
		logrus.SetLevel(l)
	}
}

// Load loads the configuration file into C
func Load(fp string) error {
	var err error
	var c []byte

	if c, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	if err = yaml.Unmarshal(c, &C); err != nil {
		return err
	}
	return C.Parse()
}
