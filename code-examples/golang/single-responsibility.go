package golang

import (
	"errors"
	"fmt"
	"os"
)

// in file authorizer.go

type Authorizer interface {
	ConfirmAuthorized(user string) error
}

type authorizer struct {
	admin string
}

func NewAuthorizer(admin string) Authorizer {
	return &authorizer{admin}
}

func (a *authorizer) ConfirmAuthorized(user string) error {
	if user != a.admin {
		return errors.New("unauthorized")
	}
	return nil
}

// in file calculation.go

type Calculation interface {
	Calculate(input int) int
}

type Multiplier struct {
	factor int
}

func NewMultiplier(x int) Calculation {
	return &Multiplier{factor: x}
}

func (m *Multiplier) Calculate(input int) int {
	return m.factor * input
}

type Logger interface {
	Log(s string)
}

type logger struct {
}

func (l *logger) Log(s string) {
	fmt.Fprint(os.Stderr, s)
}

type OutputFormatter interface {
	Format(value int) string
}

type outputFormatter struct {
	theFormat string
	units     string
}

func NewOutputFormatter(template string, units string) OutputFormatter {
	if template == "" {
		template = "The product is %d %s"
	}
	if units == "" {
		units = "units"
	}
	return &outputFormatter{
		theFormat: template,
		units:     units,
	}
}

func (f *outputFormatter) Format(value int) string {
	return fmt.Sprintf(f.theFormat, value, f.units)
}

type Service struct {
	authorizer  Authorizer
	calculation Calculation
	logger      Logger
	formatter   OutputFormatter
}

func NewService(authorizer Authorizer, calculation Calculation, logger Logger, formatter OutputFormatter) *Service {
	return &Service{
		authorizer:  authorizer,
		calculation: calculation,
		logger:      logger,
		formatter:   formatter,
	}
}

func (s *Service) Execute(user string, input int) (string, error) {
	if e := s.authorizer.ConfirmAuthorized(user); e != nil {
		return "", e
	}
	s.logger.Log(fmt.Sprintf("executing %d", input))
	v := s.calculation.Calculate(input)
	out := s.formatter.Format(v)
	return out, nil
}

func RunService() {
	auth := NewAuthorizer("pat")
	calc := NewMultiplier(3)
	log := &logger{}
	formatter := NewOutputFormatter("", "things")
	svc := NewService(auth, calc, log, formatter)
	out, _ := svc.Execute("pat", 8)
	fmt.Println(out)
}
